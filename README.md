# timer
A simple countdown timer written in Go

<a href="https://asciinema.org/a/tAtuba3KGefNfPBTqpbexl09j" target="_blank"><img src="https://asciinema.org/a/tAtuba3KGefNfPBTqpbexl09j.png" /></a>


## Install

```go get github.com/zealotree/timer/```

## Usage

```
timer 2s 
timer 1h 
timer 525600m
timer 1h30m10s
```



## Play a beep after your timer

```timer 2s ; mpv --really-quiet beep.ogg ```

### Make some presets using aliases

*fish shell 2.6*
```
abbr -a 'greentea' 'timer 1m20s ; mpv --really-quiet teatimer.ogg'

```
