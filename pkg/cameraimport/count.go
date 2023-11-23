package cameraimport

import (
  "io/fs"
  "path/filepath"
)

func (r *MediaRepo) CountMediaFiles(mediaPath string) error {
  // count files by type
  filepath.WalkDir(mediaPath, func(path string, d fs.DirEntry, err error) error {
    // skip directories
    if d.IsDir() { return nil; }

    // count raws
    if CompareExtensions(path, rawFormats[:], true) {
      r.rawFiles++;
    }
    // count rasters
    if CompareExtensions(path, rasterFormats[:], true) {
      r.rasterFiles++;
    }
    // count movies
    if CompareExtensions(path, videoFormats[:], true) {
      r.videoFiles++;
    }

    return nil;
  });

  return nil;
}

func (r *MediaRepo) Rasters() int {
  return r.rasterFiles;
}

func (r *MediaRepo) Raws() int {
  return r.rawFiles;
}

func (r *MediaRepo) Videos() int {
  return r.videoFiles;
}
