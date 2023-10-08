# SpeedTester-Go
SpeedTester in Go

| Parameter   | Description                                                                                                 |
|-------------|-------------------------------------------------------------------------------------------------------------|
| `-url`      | download url of the target for the measurement e.g. `https://download.com/video.mp4`                        |
| `-interval` | interval (in seconds) for periodic measurement. Use `0` to execute just a single (non-periodic) measurement |
| `-size`     | amount of download bytes used for measurements.                                                             |


## Feature Backlog
- [x] speedtests can run periodically
- [ ] amount of download bytes is configurable


## Lessons learned
* go modules are not necessary for structuring the code base. there are also packages
* visibility of types and functions are controlled by the case of the first letter
* go has a compact code base
* the online docs are good
* `defer` keyword is definitely a cool thing 