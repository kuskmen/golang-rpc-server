package yamq

import (
	"testing"
)

func TestWarehouse_NewMessageQueue(t *testing.T) {
	type args struct {
		name  string
		reply *MessageQueue
	}
	tests := []struct {
		name    string
		w       *Warehouse
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.w.NewMessageQueue(tt.args.name, tt.args.reply); (err != nil) != tt.wantErr {
				t.Errorf("Warehouse.NewMessageQueue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMessageQueue_Send(t *testing.T) {
	type args struct {
		message []byte
		reply   *MessageQueue
	}
	tests := []struct {
		name    string
		queue   *MessageQueue
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.queue.Send(tt.args.message, tt.args.reply); (err != nil) != tt.wantErr {
				t.Errorf("MessageQueue.Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMessageQueue_Subscribe(t *testing.T) {
	type args struct {
		subscriber chan []byte
		reply      *MessageQueue
	}
	tests := []struct {
		name    string
		queue   *MessageQueue
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.queue.Subscribe(tt.args.subscriber, tt.args.reply); (err != nil) != tt.wantErr {
				t.Errorf("MessageQueue.Subscribe() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMessageQueue_Unsubscribe(t *testing.T) {
	type args struct {
		subscriber chan []byte
		reply      *MessageQueue
	}
	tests := []struct {
		name    string
		queue   *MessageQueue
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.queue.Unsubscribe(tt.args.subscriber, tt.args.reply); (err != nil) != tt.wantErr {
				t.Errorf("MessageQueue.Unsubscribe() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
