// アプリケーション層用のデータ受け渡しオブジェクト定義
// プレゼンテーション層はドメインモデルではなくこちらを参照する
// メリット
// 1 プレゼンテーション由来のロジックをドメインモデルにしみださせないことが目的
// 2 一覧画面のような複雑なクエリに対応したモデルを置ける(CQRS)
package dto
