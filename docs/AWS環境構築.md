## VPCの作成

- AZ１つ
- パブリックサブネットの数 1
- プライベートサブネットの数 0
- NAT ゲートウェイ 0  
- VPC エンドポイント なし  
https://ap-northeast-1.console.aws.amazon.com/vpcconsole/home?region=ap-northeast-1#VpcDetails:VpcId=vpc-09fe18374bdee2709



## ECRの作成

https://ap-northeast-1.console.aws.amazon.com/ecr/repositories/private/851725409176/homekitchenadventure?region=ap-northeast-1


## ECSの作成

1. Fargate起動モデルでクラスタを作成  
無料枠はあるらしいが、無駄に通信を発生させないように注意する

1. タスク定義の作成  
「Fargate」を起動タイプとして選択します。  
容量は、CPU:1 RAM:2GB

1. サービスの作成  
起動タイプ FARGATE を選択する  
アプリケーションタイプはサービスを選び、ファミリーに先ほど作ったタスクを選択

1.












### 参考資料

- https://dev.classmethod.jp/articles/ecs-fargate-entry-new-console/
- https://qiita.com/bonjiko/items/485adffcf624dec1d7b1#ecs