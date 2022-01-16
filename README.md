# kakemoti
kakemoti is a simple tool that executes workflows defined in the [Amazon States Language](https://states-language.net/). It breaks up large scripts into pieces to improve readability, observability and serviceability.

# TODO
- [x] Top-level fields
  - [x] States
  - [x] StartAt
  - [x] Comment
  - [x] Version
  - [x] TimeoutSeconds
- [ ] States
  - [x] Pass State
  - [x] Task State
  - [ ] Choice State
    - [x] Boolean expression
    - [ ] Data-test expression
      - [ ] StringEquals
      - [ ] StringEqualsPath
      - [ ] StringLessThan
      - [ ] StringLessThanPath
      - [ ] StringGreaterThan
      - [ ] StringGreaterThanPath
      - [ ] StringLessThanEquals
      - [ ] StringLessThanEqualsPath
      - [ ] StringGreaterThanEquals
      - [ ] StringGreaterThanEqualsPath
      - [ ] StringMatches
      - [ ] NumericEquals
      - [ ] NumericEqualsPath
      - [ ] NumericLessThan
      - [ ] NumericLessThanPath
      - [ ] NumericGreaterThan
      - [ ] NumericGreaterThanPath
      - [ ] NumericLessThanEquals
      - [ ] NumericLessThanEqualsPath
      - [ ] NumericGreaterThanEquals
      - [ ] NumericGreaterThanEqualsPath
      - [x] BooleanEquals
      - [x] BooleanEqualsPath
      - [ ] TimestampEquals
      - [ ] TimestampEqualsPath
      - [ ] TimestampLessThan
      - [ ] TimestampLessThanPath
      - [ ] TimestampGreaterThan
      - [ ] TimestampGreaterThanPath
      - [ ] TimestampLessThanEquals
      - [ ] TimestampLessThanEqualsPath
      - [ ] TimestampGreaterThanEquals
      - [ ] TimestampGreaterThanEqualsPath
      - [x] IsNull
      - [x] IsPresent
      - [x] IsNumeric
      - [x] IsString
      - [x] IsBoolean
      - [x] IsTimestamp
  - [x] Wait State
  - [x] Succeed State
  - [x] Fail State
  - [x] Parallel State
  - [ ] Map State
    - [ ] Map State input/output processing
    - [ ] Map State concurrency
    - [ ] Map State Iterator definition
- [x] Transitions
- [ ] Timestamps
- [ ] Data
- [x] The Context Object
- [x] Paths
- [x] Reference Paths
- [x] Payload Template
- [x] Intrinsic Functions
  - [x] States.Format
  - [x] States.StringToJson
  - [x] States.JsonToString
  - [x] States.Array
- [x] Input and Output Processing
  - [x] InputPath
  - [x] Parameters
  - [x] ResultSelector
  - [x] ResultPath
  - [x] OutputPath
- [ ] Errors
  - [x] States.ALL
  - [ ] States.HeartbeatTimeout
  - [ ] States.Timeout
  - [x] States.TaskFailed
  - [ ] States.Permissions
  - [ ] States.ResultPathMatchFailure
  - [ ] States.ParameterPathFailure
  - [ ] States.BranchFailed
  - [ ] States.NoChoiceMatched
  - [ ] States.IntrinsicFailure
