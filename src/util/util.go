package util

import (
    "runtime"
    "errors"
    "log"
    "os"
    "os/user"
)

func FatalErr (err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func PathExists(assetPath string) (bool, error) {
    _, err := os.Stat(assetPath)
    if err == nil {
        return true, nil
    }

    if os.IsNotExist(err) {
        return false, nil
    }
    return false, err
}

func IsLinux() (platform string, err error) {
    if runtime.GOOS == "windows" {
        return "win", nil
    } else if runtime.GOOS == "linux" {
        return "linux", nil
    } else {
        return "", errors.New("Unsupported platform detected!")
    }
}

func IsAdmin() (admin bool) {
    currentPlatform, err := IsLinux()
    if err != nil {
        return false
    }

    if currentPlatform == "linux" {

        currentUser, err := user.Current()
        if err != nil {
            return false
        }

        return currentUser.Username == "root"
    } else if currentPlatform == "win" {
        /*
        There isn't a reliable way to check for admin on win
        as there is on linux so until I find a solution
        this will return true and will handle permission errors
        and other stuff in the function that'll execute commands
        */
        return true
    }
    return false
}