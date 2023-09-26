# **Go Media Importer** #

This is a small piece of software that transfers photos from an external source (such as an SD Card from a DSLR) to a work path on your laptop,
in order to be further processed (with Darktable or similar software).

Photos are sorted by creation time, and saved in a directory tree in the destination folder with this structure:

    DESTINATION PATH
      \
      YEAR
        \
        MONTH
          \
          DAY
            \
            <file type: RAW or JPEG>

Furthermore, JPEG and RAW files are kept apart from each other.
Photo creation times are read from the EXIF metatata headers if present, otherwise the last modification time from the stat() syscall is used as timestamp.

