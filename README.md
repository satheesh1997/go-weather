# go-weather
Script to show hourly weather update

cli based:
  1. move the weather file to /usr/bin/
  2. type weather in terminal to get the weather update

Hourly Cron
  1. move the weather file /etc/con.hourly
  2. restart crontab


## Development
1. Clone the repo 
2. run go get -v github.com/martinlindhe/notify
3. go run weather.go

## Note:
  The script uses banglore as the location, you can change the 
lat and long and build again
