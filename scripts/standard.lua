-- Fly around using standard drone controls on the keyboard
-- WASD: rotate left, rotate right, up, and down
-- Arrow keys: forward, backward, left, and right
-- Escape: exit script
-- Spacebar: reset movement in all directions (this is an important one to remember)
-- X: take off
-- Z: land
-- O: reduce speed
-- P: increase speed
-- Q: left flip
-- E: right flip
-- F: front flip
-- R: back flip
-- B: bounce up and down near the ground

speed = 50
log("set initial speed to " .. speed)

log("start video stream")
startvideo()

repeat
    lastkey = getlastkeypress()

    if lastkey == "a" then
        rotateleft(speed)
    elseif lastkey == "d" then
        rotateright(speed)
    elseif lastkey == "w" then
        up(speed)
    elseif lastkey == "s" then
        down(speed)
    elseif lastkey == "arrowup" then
        up(speed)
    elseif lastkey == "arrowdown" then
        down(speed)
    elseif lastkey == "arrowleft" then
        left(speed)
    elseif lastkey == "arrowright" then
        right(speed)
    elseif lastkey == "q" then
        leftflip()
    elseif lastkey == "e" then
        rightflip()
    elseif lastkey == "r" then
        backflip()
    elseif lastkey == "f" then
        frontflip()
    elseif lastkey == "b" then
        bounce()
    elseif lastkey == "o" then
        speed = speed - 5
        if speed < 0 then
            speed = 0
        end
        log("reduce speed to " .. speed)
    elseif lastkey == "p" then
        speed = speed + 5
        if speed > 100 then
            speed = 100
        end
        log("increase speed to " .. speed)
    elseif lastkey == "z" then
        log("initiating landing")
        land()
        log("landed")
    elseif lastkey == "x" then
        log("initiating take off")
        takeoff()
        log("hovering")
    elseif lastkey == nil or lastkey == "" then
        -- do nothing
    else
        stop()
    end
until lastkey == "escape"

log("execution terminated")