workflow "CI" {
  on = "push"
  resolves = ["test-unit", "lint"]
}

action "test-unit" {
  uses = "cedrickring/golang-action@1.3.0"
  args = "make test-unit"
  env = {
    GO111MODULE = "on"
  }
}

action "lint" {
  uses = "./.github/action-lint"
  args = "run"
}
