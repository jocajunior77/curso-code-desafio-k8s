package main

import "testing"

func TestSum(t *testing.T ) {
    msg := greeting("Code.education Rocks!")

    if msg != "<b>Code.education Rocks!</b>" {
        t.Errorf("Saudacao invalida")
    }

}