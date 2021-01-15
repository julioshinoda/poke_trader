package trade

import (
	"os"
	"testing"
)

func Test_tradeManager_isFairTrade(t *testing.T) {
	type args struct {
		indexOne int
		indexTwo int
	}
	tests := []struct {
		name string
		tm   tradeManager
		args args
		want bool
	}{
		{
			name: "[T0] fair trade",
			tm:   tradeManager{},
			args: args{
				indexOne: 60,
				indexTwo: 50,
			},
			want: true,
		},
		{
			name: "[T1] fair trade",
			tm:   tradeManager{},
			args: args{
				indexOne: 45,
				indexTwo: 50,
			},
			want: true,
		},
		{
			name: "[T2] unfair trade",
			tm:   tradeManager{},
			args: args{
				indexOne: 140,
				indexTwo: 50,
			},
			want: false,
		},
	}
	os.Setenv("FAIR_INDEX", "20")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm := tradeManager{}
			if got := tm.isFairTrade(tt.args.indexOne, tt.args.indexTwo); got != tt.want {
				t.Errorf("tradeManager.isFairTrade() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tradeManager_TradeCalculator(t *testing.T) {
	type fields struct {
		Repo Repository
	}
	type args struct {
		trade Trade
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "[T0] error on numbers of pokemon",
			args: args{
				trade: Trade{
					FirstTrainerList: []*Pokemon{
						&Pokemon{Name: ""},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "[T1] error on numbers of pokemon",
			args: args{
				trade: Trade{
					FirstTrainerList: []*Pokemon{
						&Pokemon{Name: ""},
						&Pokemon{Name: ""},
						&Pokemon{Name: ""},
						&Pokemon{Name: ""},
					},
					SecondTrainerList: []*Pokemon{
						&Pokemon{Name: ""},
						&Pokemon{Name: ""},
						&Pokemon{Name: ""},
						&Pokemon{Name: ""},
						&Pokemon{Name: ""},
						&Pokemon{Name: ""},
						&Pokemon{Name: ""},
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm := tradeManager{
				Repo: tt.fields.Repo,
			}
			got, err := tm.TradeCalculator(tt.args.trade)
			if (err != nil) != tt.wantErr {
				t.Errorf("tradeManager.TradeCalculator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("tradeManager.TradeCalculator() = %v, want %v", got, tt.want)
			}
		})
	}
}
