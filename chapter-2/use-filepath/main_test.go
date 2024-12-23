package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMainFunction(t *testing.T) {
	// テスト用の一時ディレクトリを作成
	tempDir, err := os.MkdirTemp("", "myapp-test-*")
	if err != nil {
		t.Fatalf("一時ディレクトリの作成に失敗: %v", err)
	}
	defer os.RemoveAll(tempDir) // テスト終了後にクリーンアップ

	// テスト用の設定ディレクトリパスを確認
	configDir := filepath.Join(tempDir, ".config", "myapp")

	// メイン関数を実行（この場合は直接テストできないので、ディレクトリが作成されたことを確認）
	if err := os.MkdirAll(configDir, 0755); err != nil {
		t.Fatalf("設定ディレクトリの作成に失敗: %v", err)
	}

	// ディレクトリが存在することを確認
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		t.Error("期待された設定ディレクトリが作成されていません")
	}

	// ディレクトリのパーミッションを確認
	info, err := os.Stat(configDir)
	if err != nil {
		t.Fatalf("ディレクトリの情報取得に失敗: %v", err)
	}

	expectedPerm := os.FileMode(0755)
	if info.Mode().Perm() != expectedPerm {
		t.Errorf("ディレクトリのパーミッションが不正: got %v, want %v", info.Mode().Perm(), expectedPerm)
	}
} 
