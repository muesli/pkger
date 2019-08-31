package memfs

import (
	"testing"

	"github.com/markbates/pkger/here"
	"github.com/stretchr/testify/require"
)

func Test_Stat(t *testing.T) {
	r := require.New(t)

	fs, err := New(here.Info{})
	r.NoError(err)
	_, err = fs.Stat("/i.dont.exist")
	r.Error(err)

	f, err := fs.Create("/i.exist")
	r.NoError(err)
	r.NoError(f.Close())

	fi, err := fs.Stat("/i.exist")
	r.NoError(err)
	r.Equal("/i.exist", fi.Name())
}
