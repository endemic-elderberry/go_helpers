package panic_if


func PanicIf(err error) {
  if err != nil {
    panic(err)
  }
}
