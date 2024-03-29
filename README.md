﻿# primestrike
# アプリケーションの概要
相手のボールの数字を自身のボールをぶつけることで素因数分解していくゲームをGoで作成した。マウス入力で壁を作成することができ、その壁にボールを反射させることによって自分の球の数字がn×自分の球の数字(nは発射時の数字)となっていく。

## 実装内容
* マウス入力で壁を作成
* 球のアニメーション
* 壁や球同士の当たり判定
* 制限時間
* 球のクールタイム
* スコア
* 球の発射

# ディレクトリ構成
<pre>
.
└─myproject
        ball.go
        enemy.go
        game.go
        go.mod
        go.sum
        line.go
        main.go
        open.png
        README.md
        shootnum.go
</pre>
* ball.go
    自身のボールの構造体と関数を記述している。
* enemy.go
    標的のボールの構造体と関数を記述している。
* game.go
    ゲームの進行の流れとそれに必要な関数を記述している。
* line.go
    壁の構造体と関数を記述している。
* main.go
    main関数を記述している。
* shootnum.go
    数字のクールタイムに関する関数を記述している。

## 構造体とフィールド

* ball
    * x (float32)
        ボールの中心のx座標
    * y (float32)
        ボールの中心のy座標
    * vx (float32)
        ボールの速度ベクトルのx方向
    * vy (float32)
        ボールの速度ベクトルのy方向
    * speed (float32)
        ボールのスピード
    * number (int)
        ボールの数字
    * num (int)
        ボールの選択した数字
    * color (color.NRGBA)
        ボールの色
    * shoot (bool)
        ボールが発射しているか
* enemy
    * x (float32)
        標的の中心のx座標
    * y (float32)
        標的の中心のy座標
    * r (float32)
        標的の半径
    * color (color.NRGBA)
        標的の色
    * number (int)
        標的の数字
* game
    * status (status(int))
        ゲームの進行と流れの状態
    * ball (ball)
        自身のボール
    * enemy (enemy)
        標的
    * timer (int)
        時間
    * count (int)
        発射した球数
    * font (font.Face)
        テキストのフォント
    * shootnum ([]*shootnum)
        打てる球
    * lines ([]*line)
        壁のスライス
    * lineG (line)
        入力中の壁
    * width (int)
        横幅
    * height (int)
        縦幅
* line
    * press_x (float32)
        マウスで押されたx座標
    * press_y (float32)
        マウスで押されたy座標
    * releas_x (float32)
        マウスで離されたx座標
    * releas_y (float32)
        マウスで離されたy座標
    * strokeWidth (float32)
        壁の太さ
    * color (color.NRGBA)
        壁の色
* shootnum
    * number (int)
        素数の数字
    * flag (bool)
        発射可能か
    * time (int)
        発射してからの時間
