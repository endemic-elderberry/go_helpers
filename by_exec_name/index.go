package by_exec_name

import "os"
import "path"
import "regexp"

type ExecStruct {
  ExecMap map[string]func([]string)
}
var globalExec *ExecStruct;
func makeGlobal() {
  if globalExec == nil {
    globalExec = &ExecStruct{
      ExecMap: make(map[string]func([]string))
    }
  }
}

func (*ExecStruct es) Configure(name string, callback func([]string)) error {
  _, exists := es.ExecMap[name]
  if exists {
    return errors.New(name + " already exists")
  }
  es.ExecMap[name] = callback;
  return nil
}
func Configure(name string, callback func([]string)) error {
  makeGlobal()
  return globalExec.Configure(name, callback)
}
func (*ExecStruct es) Execute() {
  execName := os.Args[0]
  subArgs := os.Args[1:]
  execCall, ok := ex[execName]
  if ok {
    execCall(subArgs)
    return;
  }
  panic(execName + " not found")
}
func Execute() {
  makeGlobal()
  return globalExec.Execute()
}

func GenerateMainFromDirectory(urlPrepend string, directoryLocation string) (string, error) {
  dir, err := os.Open(directoryLocation)
  if err != nil {
    return "", err
  }
  files, err := dir.Readdirnames(-1)
  if err != nil {
    return "", err
  }
  err = dir.Close()
  if err != nil {
    return "", err
  }
  regTop := regexp.MustCompile(`[\"]`)
  
  clearFiles := make([]string, 0)
  for i := 0 ; i < len(files) ; i++ {
    if regTop.MatchString(files[i]) {
      continue
    }
    clearFiles = append(clearFiles, files[i])
  }
  
  urlFiles := make([]string, len(clearFiles))
  for i := 0 ; i < len(clearFiles) ; i++ {
    urlFiles[i] = path.Join(urlPrepend, clearFiles[i])
  }
  sOut := "package main\n\n"
  for i := 0 ; i < len(urlFiles) ; i++ {
    sOut += "import \"" + urlFiles[i] + "\""
  }
  sOut += `import "github.com/endemic-elderbery/GoHelpers/by_exec_name"

  func main() {
`
  for i := 0 ; i < len(files) ; i++ {
    sOut += "    "
    sOut += "by_exec_name.Configure(\"" + files[i] + "\", " + files[i] + ".Main);"
    sOut += "\n"
  }
    
  sOut += `
    by_exec_name.Execute();
  }
  
`
  return sOut, nil
}
