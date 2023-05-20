package hello

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"testing"
)

func TestHelloController_SayHello(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockHelloRepositoryInterface(ctrl)

	type fields struct {
		repo HelloRepositoryInterface
	}
	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
		runFn   func()
	}{
		{
			name:   "Example  Failed Test",
			fields: fields{repo: mockRepo},
			args:   args{ctx: echo.New().NewContext(nil, nil)},
			runFn: func() {
				mockRepo.EXPECT().GetHelloMessage().Return("", errors.New("err"))
			},
			wantErr: true,
		},
		{
			name:   "Example Success Test",
			fields: fields{repo: mockRepo},
			args:   args{ctx: echo.New().NewContext(nil, nil)},
			runFn: func() {
				mockRepo.EXPECT().GetHelloMessage().Return("expected", nil)
			},
			want:    "expected",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.runFn()
			c := HelloController{
				repo: tt.fields.repo,
			}
			got, err := c.SayHello(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("SayHello() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SayHello() got = %v, want %v", got, tt.want)
			}

		})
	}
}
