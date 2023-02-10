# ftdata
Unofficial CLI tool to retrieve device data and test results from the Focustronic API.

## Features
* List all devices within account to get device IDs
* Supports Alkatronic, Dosetronic, and Mastertronic
* Can output in table or CSV format
* Outputs to terminal or to file 

## Examples

1. Set credentials in your shell:
```shell
export FOCUSTRONIC_USERNAME='user here'
export FOCUSTRONIC_PASSWORD='password here'
```
or for Windows:
```shell
$env:FOCUSTRONIC_USERNAME = 'user here'
$env:FOCUSTRONIC_PASSWORD = 'password here'
```
2. List all the devices in your account:
```shell
$ ftdata list-devices
Device Name: Alkatronic, Device ID: 001
Device Name: Dosetronic, Device ID: 001
Device Name: Mastertronic, Device ID: 001
```
3. Pull Alkatronic records and display in the default table format:

```shell
$ ftdata get -t alkatronic -i 001
KhValue  Solution Added(mL)  Time
6.67     0.00                2023-02-03T20:09:44-06:00
6.67     0.00                2023-02-04T00:09:44-06:00
6.67     0.00                2023-02-04T04:09:34-06:00
6.67     0.00                2023-02-04T08:09:17-06:00
6.77     0.00                2023-02-04T12:09:28-06:00
6.67     0.00                2023-02-04T16:09:35-06:00
6.57     0.00                2023-02-04T20:09:30-06:00
```
4. Pull mastertronic records and output as CSV:
```shell
$ ftdata get -t mastertronic -i 001 -o csv
Parameter,Value,Time
ca,410.00,2023-02-05T09:33:22-06:00
no3,4.90,2023-02-05T10:23:52-06:00
oli,0.00,2023-02-07T23:03:51-06:00
no3,4.04,2023-02-08T00:21:51-06:00
ca,397.00,2023-02-08T05:24:41-06:00
```