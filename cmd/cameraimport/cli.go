package cameraimport

import (
  "fmt"
  "os"
  "log"

  "github.com/spf13/cobra"
  t "github.com/jedib0t/go-pretty/v6/text"

  "github.com/mcaimi/cameraimport/pkg/cameraimport"
)

var (
  configFile string
  rootCommand = &cobra.Command{
    Use: "cameraimport",
    Short: "Manage media files from Cameras and storage devices",
    Long: "Manage media files from Cameras and storage devices.",
  }

  importCommand = &cobra.Command{
    Use: "import",
    Short: "Import media files from Cameras or SD Cards",
    Long: "Import and categorize media files stored in Cameras or SD Cards.",
    Aliases: []string{"i", "add"},
    Run: func(cmd *cobra.Command, args []string) { 
      sourceDir, err := cmd.Flags().GetString("sourcedir");
      if err != nil {
        log.Fatalf("Import Command: Fatal Error %s\n", t.FgRed.Sprintf("%v", err));
      }
      destDir, err := cmd.Flags().GetString("destinationdir");
      if err != nil {
        log.Fatalf("Import Command: Fatal Error %s\n", t.FgRed.Sprintf("%v", err));
      }

      if ! cameraimport.StringNotZeroLen(sourceDir) {
        log.Fatalf("Import Command: Sanity Check Error %s\n", t.FgRed.Sprintf("%v", err));
      }
      if ! cameraimport.StringNotZeroLen(destDir) {
        log.Fatalf("Import Command: Sanity Check Error %s\n", t.FgRed.Sprintf("%v", err));
      }

      // walk directory and count files
      var mediaRepo cameraimport.MediaRepo;
      mediaRepo.CountMediaFiles(sourceDir);
      mediaRepo.ImportMediaFiles(sourceDir, destDir);
      fmt.Printf("-> Transferred a total of [%s] new media files - [%s] RAWs [%s] Rasters [%s] Videos.\n", t.FgGreen.Sprintf("%d", mediaRepo.TransferredFiles()), t.FgGreen.Sprintf("%d", mediaRepo.Raws()), t.FgGreen.Sprintf("%d", mediaRepo.Rasters()), t.FgGreen.Sprintf("%d", mediaRepo.Videos()));
      os.Exit(0);
    },
  }

  countCommand = &cobra.Command{
    Use: "count",
    Short: "Media file count.",
    Long: "Count how many files will be copied over.",
    Aliases: []string{"c", "scan"},
    Run: func(cmd *cobra.Command, args []string) { 
      sourceDir, err := cmd.Flags().GetString("sourcedir");
      if err != nil {
        log.Fatalf("Count Command: Fatal Error %s\n", t.FgRed.Sprintf("%v", err));
      }
      if ! cameraimport.StringNotZeroLen(sourceDir) {
        log.Fatalf("Count Command: Fatal Error %s\n", t.FgRed.Sprintf("%v", err));
      }

      // walk directory and count files
      var mediaRepo cameraimport.MediaRepo;
      mediaRepo.CountMediaFiles(sourceDir);

      // display
      fmt.Printf("Media Repository Scan Completed: Found [%d] RAW Files, [%d] Raster Files and [%d] Video Files.\n", mediaRepo.Raws(), mediaRepo.Rasters(), mediaRepo.Videos());
      os.Exit(0);
    },
  }
)

func init() {
  importCommand.Flags().StringP("sourcedir", "s", "", "Source Folder");
  importCommand.Flags().StringP("destinationdir", "d", "", "Destination Folder");

  countCommand.Flags().StringP("sourcedir", "s", "", "Source Folder");

  rootCommand.AddCommand(importCommand);
  rootCommand.AddCommand(countCommand);
}

func Execute() {
  if err := rootCommand.Execute(); err != nil {
    fmt.Fprintf(os.Stderr, err.Error());
    os.Exit(1);
  }
}

