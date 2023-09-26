package cameraimport

import (
  "os"
  "fmt"
  "io"
  "io/fs"
  "path"
  "path/filepath"

  "github.com/mcaimi/goexif/exif"
  t "github.com/jedib0t/go-pretty/v6/text"
)

func copy(src, dst string) (int64, error) {
  // does source file exist?
  _, err := os.Stat(src)
  if err != nil {
    return 0, err
  }

  source, err := os.Open(src)
  if err != nil {
    return 0, err
  }
  defer source.Close()

  // check whether the destination file is already there
  destStats, err := os.Stat(dst)
  if err == nil && destStats.Mode().IsRegular() {
    return 0, fmt.Errorf(t.FgRed.Sprintf("import.copy(): Skipping file %s, it already exists.", dst));
  }

  destination, err := os.Create(dst)
  if err != nil {
    return 0, err
  }
  defer destination.Close()

  // copy file over
  nBytes, err := io.Copy(destination, source)
  return nBytes, nil
}

func copyLoop(fPath string, dPath string, fType string, validFormats []string, copyCounter *int) error {
  if CompareExtensions(fPath, validFormats, true) {
    var outPath string;

    // examine EXIF metadata
    fmt.Printf("Media File [%s]... \t", t.FgBlue.Sprintf(path.Base(fPath)));
    mediaFile, e := os.Open(fPath);
    if e != nil {
      fmt.Println(e);
      return e;
    }
    defer mediaFile.Close();

    // parse metadata
    exifData, e := exif.Decode(mediaFile);
    if e != nil {
      // no exif found. use file creation date as timestamp
      sourceStats, err := os.Stat(fPath);
      if err == nil {
        fileTimeStamp := sourceStats.ModTime();
        outPath = fmt.Sprintf("%s/%d/%d/%d/%s", dPath, fileTimeStamp.Year(), fileTimeStamp.Month(), fileTimeStamp.Day(), fType);
      }
    } else {
      exifTimeStamp, _ := exifData.DateTime();
      outPath = fmt.Sprintf("%s/%d/%d/%d/%s", dPath, exifTimeStamp.Year(), exifTimeStamp.Month(), exifTimeStamp.Day(), fType);
    }

    // prepare destination
    if err := os.MkdirAll(outPath, 0700); err != nil {
      return err;
    }

    // copy file
    bytesCopied, err := copy(fPath, fmt.Sprintf("%s/%s", outPath, path.Base(fPath))); 
    if err != nil {
      fmt.Println(err);
    } else {
      *copyCounter++;
      fmt.Printf("Copy completed: [%s] bytes transferred to [%s]\n", t.FgGreen.Sprintf("%d", bytesCopied), t.FgYellow.Sprintf("%s", outPath));
    }
  }

  return nil;
}

func (r *MediaRepo) ImportMediaFiles(sourcePath, destPath string) error {
  // walk over files
  filepath.WalkDir(sourcePath, func(fPath string, d fs.DirEntry, err error) error {
    if d.IsDir() { return nil; }

    if err := copyLoop(fPath, destPath, "rasters", rasterFormats[:], &r.transferredFiles); err != nil {
      return err
    }

    if err := copyLoop(fPath, destPath, "raws", rawFormats[:], &r.transferredFiles); err != nil {
      return err
    }

    return nil;
  });

  return nil;
}

func (r *MediaRepo) TransferredFiles() int {
  return r.transferredFiles;
}
