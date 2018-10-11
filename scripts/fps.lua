speed = 20
log("set initial speed to " .. speed)

log("start video stream")
startvideo()

repeat
    lastkey = getlastkeypress()

    if lastkey == "A" then
        left(speed)
    elseif lastkey == "D" then
        right(speed)
    elseif lastkey == "W" then
        forward(speed)
    elseif lastkey == "S" then
        backward(speed)
    elseif lastkey == "arrowup" then
        up(speed)
    elseif lastkey == "arrowdown" then
        down(speed)
    elseif lastkey == "arrowleft" then
        rotateleft(speed)
    elseif lastkey == "arrowright" then
        rotateright(speed)
    elseif lastkey == "Q" then
        leftflip()
    elseif lastkey == "E" then
        rightflip()
    elseif lastkey == "R" then
        backflip()
    elseif lastkey == "F" then
        frontflip()
    elseif lastkey == "B" then
        bounce()
    elseif lastkey == "O" then
        speed = speed - 5
        if speed < 0 then
            speed = 0
        end
        log("reduce speed to " .. speed)
    elseif lastkey == "P" then
        speed = speed + 5
        if speed > 100 then
            speed = 100
        end
        log("increase speed to " .. speed)
    elseif lastkey == "Z" then
        log("initiating landing")
        land()
        log("landed")
    elseif lastkey == "X" then
        log("initiating take off")
        takeoff()
        log("hovering")
    else
        stop()
    end
until lastkey == "escape"

log("execution terminated")