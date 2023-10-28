package dark

import (
	"errors"
	"testing"
)

func TestMust(t *testing.T) {
	type args struct {
		r   interface{}
		err error
	}
	tests := []struct {
		name      string
		args      args
		want      interface{}
		wantPanic bool
	}{
		{
			name: "without panic when error is nil, and return value equals to input",
			args: args{
				r:   "hello",
				err: nil,
			},
			want:      "hello",
			wantPanic: false,
		},
		{
			name: "with panic when error is not nil",
			args: args{
				r:   nil,
				err: errors.New("some error"),
			},
			want:      nil,
			wantPanic: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					if !test.wantPanic {
						t.Errorf("unexpected panic: %v", r)
					}
				}
			}()

			got := Must(test.args.r, test.args.err)
			if test.wantPanic {
				t.Errorf("expected panic, but got none")
			}
			if got != test.want {
				t.Errorf("expected %v, got %v", test.want, got)
			}
		})
	}
}

func TestTry(t *testing.T) {
	type args struct {
		f func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "no panic",
			args: args{
				f: func() {
					// do nothing
				},
			},
			wantErr: false,
		},
		{
			name: "panic",
			args: args{
				f: func() {
					panic("some error")
				},
			},
			wantErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := func(err interface{}) {
				if !test.wantErr {
					t.Errorf("unexpected error: %v", err)
				}
			}
			Try(test.args.f)(g)
		})
	}

	t.Run("set variable in catch", func(t *testing.T) {
		x := 0
		Try(func() {
			x = 1
			panic(errors.New("some error"))
		})(func(_ any) {
			x = 10
		})
		if x != 10 {
			t.Errorf("expected x to be 10, got %v", x)
		}
	})

	t.Run("set variable if no panic", func(t *testing.T) {
		x := 0
		Try(func() {
			x = 1
		})(func(_ any) {
			x = 10
		})
		if x != 1 {
			t.Errorf("expected x to be 1, got %v", x)
		}
	})
}