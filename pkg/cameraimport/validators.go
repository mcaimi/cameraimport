package cameraimport

import "strings"

func StringNotZeroLen(s string) bool {
  if len(s) == 0 {
    return false;
  }

  return true;
}

// Compare Extensions
func CompareExtensions(filename string, validExtensions []string, caseFold bool) bool {
  components := strings.Split(filename, ".");
  if len(components) < 2 { return false; }

  if caseFold {
    // compare (insensitive)
    for i := range(validExtensions) {
      if strings.EqualFold(components[1], validExtensions[i]) {
        return true;
      }
    }
    return false;
  } else {
    // compare (sensitive)
    for i := range(validExtensions) {
      if strings.Compare(strings.ToUpper(components[1]), strings.ToUpper(validExtensions[i])) == 0 {
        return true;
      }
    }
    return false;
  }
}
