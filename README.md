# SpeedTester-Go
SpeedTester in Go

Configuration

| Parameter   | ENV      | Description                                                                                                 |
|-------------|----------|-------------------------------------------------------------------------------------------------------------|
| `-url`      | URL      | download url of the target for the measurement e.g. `https://download.com/video.mp4`                        |
| `-interval` | INTERVAL | interval (in seconds) for periodic measurement. Use `0` to execute just a single (non-periodic) measurement |
| `-size`     | SIZE     | amount of download bytes used for measurements.                                                             |

Both command line parameters and environment variables are supported simultaneously. Env variables are individually
higher prioritized than parameters.


## Feature Backlog
- [x] speedtests can run periodically
- [x] amount of download bytes is configurable
- [x] env-based config
- [ ] url validation on startup


## Lessons learned
* go modules are not necessary for structuring the code base. there are also packages
* visibility of types and functions are controlled by the case of the first letter
* go has a compact code base
* the online docs are good
* `defer` keyword is definitely a cool thing 