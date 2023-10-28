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

func TestMustNoErr(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name      string
		args      args
		wantPanic bool
	}{
		{
			name: "without panic when error is nil, and return value equals to input",
			args: args{
				err: nil,
			},
			wantPanic: false,
		},
		{
			name: "with panic when error is not nil",
			args: args{
				err: errors.New("some error"),
			},
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

			MustNil(test.args.err)
			if test.wantPanic {
				t.Errorf("expected panic, but got none")
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
			name: "panic with string",
			args: args{
				f: func() {
					panic("some error")
				},
			},
			wantErr: true,
		},
		{
			name: "panic with number",
			args: args{
				f: func() {
					panic(1)
				},
			},
			wantErr: true,
		},
		{
			name: "panic with error",
			args: args{
				f: func() {
					panic(errors.New("some error"))
				},
			},
			wantErr: true,
		},
		{
			name: "panic with nil",
			args: args{
				f: func() {
					panic(nil)
				},
			},
			wantErr: true,
		},
		{
			name: "panic with custom type",
			args: args{
				f: func() {
					panic(struct{}{})
				},
			},
			wantErr: true,
		},
		{
			name: "panic with pointer",
			args: args{
				f: func() {
					panic(&struct{}{})
				},
			},
			wantErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := func(err error) {
				if !test.wantErr {
					t.Errorf("unexpected error: %v", err)
				}
			}
			Try(test.args.f, g)
		})
	}

	t.Run("set variable in catch", func(t *testing.T) {
		x := 0
		Try(func() {
			x = 1
			panic(errors.New("some error"))
		}, func(_ error) {
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
		}, func(_ error) {
			x = 10
		})
		if x != 1 {
			t.Errorf("expected x to be 1, got %v", x)
		}
	})
}
func TestAbortOnErr(t *testing.T) {
	tests := []struct {
		name        string
		fn          func()
		expectedErr error
	}{
		{
			name: "no error",
			fn: func() {
			},
			expectedErr: nil,
		},
		{
			name: "panic with string",
			fn: func() {
				panic("some error")
			},
			expectedErr: errors.New("Try failed: some error"),
		},
		{
			name: "panic with error",
			fn: func() {
				panic(errors.New("some error"))
			},
			expectedErr: errors.New("Try failed: some error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := AbortOnErr(test.fn)
			if err == nil && test.expectedErr != nil {
				t.Errorf("expected error, got nil")
			}

			if err != nil && test.expectedErr == nil {
				t.Errorf("expected nil, got %v", err)
			}

			if err != nil && test.expectedErr != nil {
				errMessage := err.Error()
				expectedErrMessage := test.expectedErr.Error()
				if errMessage != expectedErrMessage {
					t.Errorf("expected %v, got %v", expectedErrMessage, errMessage)
				}
			}
		})
	}

}
