load 'helpers/common'
load 'helpers/writeMockScript'
load 'helpers/run'
load 'helpers/absolutePath'


setup() {
  _common_setup
}

teardown() {
  _common_teardown
}

@test "if some-script is run it should execute it" {
  test_if_run_is_executed_with_script_name_it_should_pass_root_dir_and_parameters_to_it "some-script" "some" "--flag"
}

@test "if another-script is run it should execute it" {
  test_if_run_is_executed_with_script_name_it_should_pass_root_dir_and_parameters_to_it "another-script" "another" "parameter"
}

test_if_run_is_executed_with_script_name_it_should_pass_root_dir_and_parameters_to_it() {
  scriptName=$1
  shift
  parameters=("$@")
  scriptLocation="$testEnvDir/run"
  writeSpyScriptToLocation "$testEnvDir/run" "$scriptName"

  run runScript "$scriptName" -- "${parameters[@]}"

  assert_spy_file_in_location_has_content "$scriptLocation" "$scriptName" "$(absolutePath "$testEnvDir") ${parameters[*]}"
}

@test "if script succeeds with output it will print the script's output" {
  scriptName="some-script"
  someOutput="some-output"
  scriptLocation="$testEnvDir/run"
  writeStubScriptToLocation "$scriptLocation" "$scriptName" "0" "$someOutput"

  run runScript "$scriptName"

  assert_output "$someOutput"
}

@test "if script is requesting input it should process the input" {
  scriptName="input"
  scriptLocation="$testEnvDir/run"
  writeScriptRequestingInputToLocation "$scriptLocation" "$scriptName"
  input="some-input"

  run runScript $scriptName <<< $input

  assert_script_received_input "$scriptLocation" "$scriptName" "$input"
}

@test "if script writes to stderr it outputs stderr" {
  scriptName="error"
  error="some-error"
  scriptLocation="$testEnvDir/run"
  writeStdErrScriptToLocation "$scriptLocation" "$scriptName" "$error"

  run runScript "$scriptName"

  assert_output "$error"
}

@test "if script fails with code 1 it will fail with error code 1 as well" {
  scriptName="some-script"
  exitCode=1
  writeStubScriptToLocation "$testEnvDir/run" "$scriptName" "$exitCode" ""

  run runScript "$scriptName"

  assert_equal "$status" "$exitCode"
  assert_failure
}

@test "if script fails with code 2 it will fail with error code 2 as well" {
  scriptName="some-script"
  exitCode=2
  writeStubScriptToLocation "$testEnvDir/run" "$scriptName" "$exitCode" ""

  run runScript "$scriptName"

  assert_equal "$status" "$exitCode"
  assert_failure
}

@test "if script exits with code 0 it will succeed" {
  scriptName="some-script"
  exitCode=0
  writeStubScript "$testEnvDir/run" "$scriptName" "$exitCode" ""

  run runScript "$scriptName"

  assert_success
}