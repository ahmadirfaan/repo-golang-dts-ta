package main

import "testing"

var (
	kubus                      = Kubus{4}
	volumeSeharusnya   float64 = 64
	luasSeharusnya     float64 = 96
	kelilingSeharusnya float64 = 48
)

func TestHitungVolume(t *testing.T) {

	t.Logf("Volume: %v", volumeSeharusnya)
	if kubus.Volume() != volumeSeharusnya {
		t.Errorf("Volume = %f; want %f", kubus.Volume(), volumeSeharusnya)
	}
}

func TestHitungLuas(t *testing.T) {
	t.Logf("Luas: %v", luasSeharusnya)

	if kubus.Luas() != luasSeharusnya {
		t.Errorf("Luas = %f; want %f", kubus.Luas(), luasSeharusnya)
	}
}

func TestHitungKeliling(t *testing.T) {
	t.Logf("Keliling: %v", kelilingSeharusnya)
	if kubus.Keliling() != kelilingSeharusnya {
		t.Errorf("Keliling = %f; want %f", kubus.Keliling(), kelilingSeharusnya)
	}
}
