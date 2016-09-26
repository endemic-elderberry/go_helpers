package command_line

import "errors"
import "fmt"
import "os"
import "sort"

type Action struct{
  Name string
  TakesArguments bool
  AliasedAs string
  Usage string
  Callback func([]string)
}
type CommandRunner struct {
  Actions map[string]*Action
  Aliases map[string]string
  Colorize func(string) string
}

func MakeCommandRunner() *CommandRunner {
  commandRunner := &CommandRunner{
    Actions: make(map[string]*Action),
    Aliases: make(map[string]string),
  }
  
  commandRunner.MustConfigureAction("help", "Prints help message", func(args []string) {
    if len(args) > 0 {
      commandRunner.PrintHelpAndExit();
    }
    commandRunner.PrintHelpForActionAndExit(args[0])
  });
  commandRunner.MustConfigureAlias("h", "help")
  
  return commandRunner
}
func (cr *CommandRunner) configureActionHelper(name string, usage string, callback func([]string), takesArgs bool) error {
  
  _, existing := cr.Actions[name]
  if existing {
    return errors.New("Action already exists")
  }
  actionObject := &Action{
    Name: name,
    TakesArguments: takesArgs,
    Callback: callback,
    Usage: usage,
  }
  
  cr.Actions[name] = actionObject
  return nil
}

func (cr *CommandRunner) ConfigureAction(name string, usage string, callback func([]string)) error {
  return cr.configureActionHelper(name, usage, callback, true)
}
func noArgWrap(subCall func()) func([]string) {
  return func(throwAwayArgs []string) {
    if len(throwAwayArgs) > 1 {
      panic(errors.New("did not expect arguments"))
    }
    subCall();
  };
}

func (cr *CommandRunner) MustConfigureActionNoArgs(name string, usage string, callback func()) {
  err := cr.configureActionHelper(name, usage, noArgWrap(callback), false)
  if err != nil { panic(err) }
}
func (cr *CommandRunner) MustConfigureAction(name string, usage string, callback func([]string)) {
  err := cr.configureActionHelper(name, usage, callback, true)
  if err != nil { panic(err) }
}
func (cr *CommandRunner) ConfigureActionNoArgs(name string, usage string, callback func()) error {
  return cr.configureActionHelper(name, usage, noArgWrap(callback), false)
}
func (cr *CommandRunner) ConfigureActionUsage(name string, usage string) error {
  actionObject, existing := cr.Actions[name]
  if !existing {
    return errors.New("Action does not exists")
  }
  actionObject.Usage = usage
  return nil
}
func (cr *CommandRunner) MustConfigureAlias(alias string, action string) {
  err := cr.ConfigureAlias(alias, action)
  if err != nil { panic(err) }
}
func (cr *CommandRunner) ConfigureAlias(alias string, action string) error {
  actionObject, existing := cr.Actions[action]
  if !existing {
    return errors.New("Action does not exist")
  }
  aliasTo, existing := cr.Aliases[alias]
  if existing {
    return errors.New(alias + " already points to "+aliasTo)
  }
  cr.Aliases[alias] = action
  actionObject.AliasedAs = alias
  return nil 
}
func (cr *CommandRunner) PrintHelpForActionAndExit(action string) {
  actionObject, aok := cr.Actions[action]
  if !aok {
    aliasTo, existing := cr.Aliases[action]
    if existing {
      actionObject = cr.Actions[aliasTo]
    }
  }
  if actionObject == nil {
    os.Stderr.WriteString(action + " not found\n")
  } else {
    os.Stderr.WriteString(actionObject.Usage + "\n")
  }
  helpStr := cr.MakeHelpString()
  os.Stderr.WriteString(helpStr)
  os.Exit(1)
}
func (cr *CommandRunner) PrintHelpAndExit() {
  helpStr := cr.MakeHelpString()
  os.Stderr.WriteString(helpStr)
  os.Exit(1)
}
func (cr *CommandRunner) mustFetchAction(name string) *Action {
  actionObject, aok := cr.Actions[name]
  if !aok {
    panic("not found")
  }
  return actionObject
}
func (cr *CommandRunner) MakeHelpString() string {
  actionArgs := make([]string, 0)
  actionNoArgs := make([]string, 0)
  for k, v := range cr.Actions {
    if v.TakesArguments {
      actionArgs = append(actionArgs, k)
    } else {
      actionNoArgs = append(actionNoArgs, k)
    }
  }
  
  out := "Actions No Args\n"
  sort.Strings(actionNoArgs)
  for i := 0 ; i < len(actionNoArgs) ; i++ {
    actionObject := cr.mustFetchAction(actionNoArgs[i])
    actName := actionObject.Name
    if cr.Colorize != nil {
      actName = cr.Colorize(actName)
    }
    out += fmt.Sprintf("  %s", actName)
    if actionObject.AliasedAs != "" {
      out += fmt.Sprintf(" | %s", actionObject.AliasedAs)
    }
    out += "\n    " + actionObject.Usage + "\n"
  }
  out += "Actions With Args:\n"
  sort.Strings(actionArgs)
  for i := 0 ; i < len(actionArgs) ; i++ {
    out += "  " + actionArgs[i] + " [args...]\n"
    actionObject := cr.mustFetchAction(actionArgs[i])
    out += "    " + actionObject.Usage + "\n"
  }
  return out
}
func (cr *CommandRunner) Run(args []string) error {
  var err error;
  if len(args) < 1 {
    err = errors.New("No command given")
    return err
  }
  command := args[0]
  
  actionObject, cok := cr.Actions[command]
  if !cok {
    alias, aok := cr.Aliases[command]
    if !aok {
      err = errors.New(command + " not found")
      return err
    }
    actionObject, cok = cr.Actions[alias]
    if !cok {
      err = errors.New(alias + " is bad alias")
      return err
    }
  }
  subarguments := args[1:]
  
  actionObject.Callback(subarguments)
  
  return nil
}
func (cr *CommandRunner) MustRun(args []string) {
  err := cr.Run(args)
  if err != nil {
    os.Stderr.WriteString(err.Error() + "\n")
    cr.PrintHelpAndExit()
  }
}