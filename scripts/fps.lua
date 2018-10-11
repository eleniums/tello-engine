speed = 20
log("set initial speed to " .. speed)

log("start video stream")
startvideo()

repeat
    lastkey = getlastkeypress()

    if(lastkey == "A")
    then
        left(speed)
    else if(lastkey == "D")
        right(speed)
    else if(lastkey == "W")
        forward(speed)
    else if(lastkey == "S")
        backward(speed)
    else if(lastkey == "arrowup")
        up(speed)
    else if(lastkey == "arrowdown")
        down(speed)
    else if(lastkey == "arrowleft")
        rotateleft(speed)
    else if(lastkey == "arrowright")
        rotateright(speed)
    else if(lastkey == "Q")
        leftflip()
    else if(lastkey == "E")
        rightflip()
    else if(lastkey == "R")
        backflip()
    else if(lastkey == "F")
        frontflip()
    else if(lastkey == "B")
        bounce()
    else if(lastkey == "O")
        speed = speed - 5
        if(speed < 0)
        then
            speed = 0
        end
        log("decrease speed to " .. speed)
    else if(lastkey == "P")
        speed = speed + 5
        if(speed > 100)
        then
            speed = 100
        end
        log("increase speed to " .. speed)
    else if(lastkey == "Z")
        log("initiating landing")
        land()
        log("landed")
    else if(lastkey == "X")
        log("initiating take off")
        takeoff()
        log("hovering")
    else
        stop()
    end
until(lastkey == "escape")

log("execution terminated")