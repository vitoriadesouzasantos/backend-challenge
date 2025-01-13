package validator

import "testing"

func TestPasswordValidator(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Invalid password",
			args: args{password: ""},
		},
		{
			name: "Invalid password",
			args: args{password: "aa"},
		},
		{
			name: "Invalid password",
			args: args{password: "ab"},
		},
		{
			name: "Invalid password",
			args: args{password: "AAAbbbCc"},
		},
		{
			name: "Invalid password",
			args: args{password: "AbTp9!foo"},
		},
		{
			name: "Invalid password",
			args: args{password: "AbTp9!foA"},
		},
		{
			name: "Invalid password",
			args: args{password: "AbTp9 fok"},
		},
		{
			name: "Valid password",
			args: args{password: "AbTp9!fok"},
			want: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			if isValid := PasswordIsValid(testCase.args.password); isValid != testCase.want {
				t.Errorf("PasswordIsValid() = %v, want %v", isValid, testCase.want)
			}
		})
	}
}
