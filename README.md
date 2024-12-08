# auto-click
适用于mac的游戏自动点击脚本，可用于阴阳师、明日方舟等游戏，使用了一些手段进行防检测

## 使用
在``auto-click``文件夹下创建一个文件夹，放入需要检测的按钮截图，如`liaotupo`
<img width="627" alt="image" src="https://user-images.githubusercontent.com/39732766/213117127-1f5bcb0d-8220-4c1b-ae6e-f616c589ce35.png">


**使用mac qq自带的截图工具**截取你想要检测并点击的按钮，放置于你创建的文件夹下
注意：检测优先级为图片的字典序
<img width="387" alt="image" src="https://user-images.githubusercontent.com/39732766/213117177-13ffa5f3-c229-41b4-9ef4-044fc26c5443.png">


打开一个终端，使用 ``./auto-click``启动脚本 ，按照macos的提示赋予相关权限（不要用直接点击程序的方法打开，可能会让macos无法正常赋予权限）
<img width="285" alt="image" src="https://user-images.githubusercontent.com/39732766/213121764-ae4c8c7c-6bc8-41a6-91dc-0074689ecb3c.png">



输入截图文件夹的名称，切换到游戏界面，开刷
<img width="836" alt="image" src="https://user-images.githubusercontent.com/39732766/213117380-03ee7666-f1c9-424e-9ed2-0ed71289af99.png">


## 关于防封
首先，opencv的图片识别函数在运行过程中存在执行时间的不确定性，使每次点击的间隔时间存在随机性。
<img width="793" alt="image" src="https://user-images.githubusercontent.com/39732766/213117469-11227f93-b2e3-4306-bbd2-63b81ef28cbb.png">


其次，使用了随机函数，每次随机点击按钮内的一个点。

但是这种处理方式会造成一个问题：真人在操作的时候，每次点击的像素点并不会是真随机分布的，而使用随机函数点击，散点图则是真随机分布，非常不自然，很容易被检测到。

<img width="224" alt="iShot_2023-01-18_15 54 03" src="https://user-images.githubusercontent.com/39732766/213117628-a8d2364f-13e5-4321-8046-65e5fa849de4.png">


所以这里使用了正态分布生成随机数，接近于真人点击效果，如图是模拟250次结果：
<img width="621" alt="image" src="https://user-images.githubusercontent.com/39732766/213117804-85c26491-c889-43d7-b2e5-b923eda79d73.png">



## 稳定性
稳定操作阴阳师刷魂土一个月。

但是仍然存在被检测的可能性，使用本程序所造成的一切后果需由自己承担。

推荐挂2-3小时左右停一段时间，不要作死挂一晚上。

