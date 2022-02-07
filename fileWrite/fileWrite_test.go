package fileWrite

import "testing"

func TestFileWrite(t *testing.T) {
    result := Hello("takeru")
    want := "Hi, takeru. Welcome!"
    if result != want {
        t.Errorf("fileWrite.Hello() = %q want %q", result, want)
    }
}