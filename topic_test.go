package golangrpcserver

import (
	"reflect"
	"testing"
)

func TestNewMessageQueue(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *MessageQueue
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMessageQueue(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMessageQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessageQueue_Send(t *testing.T) {
	type args struct {
		message []byte
	}
	tests := []struct {
		name  string
		queue *MessageQueue
		args  args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.queue.Send(tt.args.message)
		})
	}
}

func TestMessageQueue_Subscribe(t *testing.T) {
	type args struct {
		subscriber chan []byte
	}
	tests := []struct {
		name  string
		queue *MessageQueue
		args  args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.queue.Subscribe(tt.args.subscriber)
		})
	}
}

func TestMessageQueue_Unsubscribe(t *testing.T) {
	type args struct {
		subscriber chan []byte
	}
	tests := []struct {
		name  string
		queue *MessageQueue
		args  args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.queue.Unsubscribe(tt.args.subscriber)
		})
	}
}
