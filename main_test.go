// テストファイルのパッケージ名は、テスト対象のファイルと同じにします。
package main

// テストに使用する標準パッケージとテスト用のフレームワークをインポートします。
import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHandler は、handler 関数をテストするためのテストケースです。
func TestHandler(t *testing.T) {
	// テスト用のHTTPリクエストを作成します。
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// テスト用のHTTPレスポンスを受け取るレコーダーを作成します。
	rr := httptest.NewRecorder()

	// テスト対象の関数 handler を呼び出します。
	handler(rr, req)

	// レスポンスのステータスコードが期待通りのものかを検証します。
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}

	// レスポンスのボディ（本文）が期待通りのものかを検証します。
	expected := "Welcome to Home Kitchen Adventure!"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}