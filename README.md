# three-colors

[![Join the chat at https://gitter.im/three-colors-/community](https://badges.gitter.im/three-colors-/community.svg)](https://gitter.im/three-colors-/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

## Instructions

Bellow is a list of links leading to an image, read this list of images and find 3 most prevalent colors in the RGB scheme in hexadecimal format (#000000 - #FFFFFF) in each image, and write the result into a CSV file in a form of url,color,color,color.

Please focus on speed and resources. The solution should be able to handle input files with more than a billion URLs, using limited resources (e.g. 1 CPU, 512MB RAM). Keep in mind that there is no limit on the execution time, but make sure you are utilizing the provided resources as much as possible at any time during the program execution.

Answer should be posted in a git repo.

## Links

<1000 links removed>
  
## Run

```bash
$ go run three-colors.go 
22:23 $ go run three-colors.go
http://i.imgur.com/TKLs9lo.jpg,#ffffff,#fefefe,#f7f7f7
http://i.imgur.com/sXhEQez.jpg,#ffffff,#000000,#fefefe
http://i.imgur.com/tyDhTll.jpg,#070705,#050503,#060604
http://i.imgur.com/puEqa4C.jpg,#908e9c,#928fa2,#92909e
http://i.imgur.com/qvnjbVw.jpg,#a9acb5,#a8abb4,#aaadb6
http://i.imgur.com/lcEUZHv.jpg,#ffffff,#c6bcbf,#444a5a
https://i.redd.it/4m5yk8gjrtzy.jpg,#010101,#020001,#030001
http://i.imgur.com/FApqk3D.jpg,#ffffff,#000000,#f3c300
http://i.imgur.com/6pnYQUv.jpg,#181a1d,#191b1e,#17191c
https://i.redd.it/nrafqoujmety.jpg,#000000,#ffffff,#ececec
http://i.imgur.com/RycdbMO.jpg,#d6cac0,#d2c6bc,#f2f9ff
https://i.redd.it/fsuv32a1cc9y.jpg,#ffffff,#fefefe,#fdfdfd
https://i.redd.it/c5pk0vnpg3ty.jpg,#6e81cd,#6e81ce,#fefefe
https://i.redd.it/1nlgrn49x7ry.jpg,#0f0f0f,#101010,#13100d
https://i.redd.it/ihczg3pmle3z.jpg,#b0b9a8,#afb8a7,#a3ac9b
https://i.redd.it/s5viyluv421z.jpg,#f5eedd,#f6efde,#7a522f
https://i.redd.it/h5t3qddvvzry.jpg,#0e0e0e,#0d0d0d,#020202
https://i.redd.it/cpkq97klre4z.jpg,#777868,#767767,#b4b4a4
https://i.redd.it/tb5pckm6e72y.jpg,#b8b6ab,#a0a19d,#a1a29e
http://i.imgur.com/pt9rmrv.jpg,#ffffff,#fcebdb,#fbeada
https://i.redd.it/n9jq5j2tatcy.jpg,#f3ffff,#f4ffff,#f3fffd
https://i.redd.it/ax4eueu0szpy.jpg,#222739,#bcbac9,#bcbac7
https://i.redd.it/cwljeujku33z.jpg,#e9e6df,#e4e1dc,#eae5de
https://i.redd.it/fyqzavufvjwy.jpg,#ddb083,#fefefe,#dbae81
https://i.redd.it/ru12lvaf7jyy.jpg,#9d9997,#9c9896,#9e9a98
https://i.redd.it/rubsfuf64lvy.jpg,#bbb4a2,#bcb5a3,#bab3a1
https://i.redd.it/xae65ypfqycy.jpg,#f4f4f4,#c83a2e,#c24541
https://i.redd.it/w5q6gldnvcuy.jpg,#f3f3f3,#19161d,#18151c
http://i.imgur.com/M3NOzLC.jpg,#ffffff,#93907d,#928f7c
https://i.redd.it/guul6eld2soy.jpg,#6b7b94,#6d7d96,#807967
https://i.imgur.com/HRgRMOw.jpg,#a4a3a8,#408a59,#40895a
https://i.redd.it/0qow2dksb34z.jpg,#ffffff,#a0caec,#fffffe
https://i.redd.it/sn7ufk9fsv4z.jpg,#ccd1d7,#cbd0d6,#cacfd5
https://i.redd.it/d8021b5i2moy.jpg,#ffffff,#010304,#020405
http://i.imgur.com/Gh2I3pq.jpg,#ceb9a8,#c7b09e,#cfbaa9
https://i.redd.it/m4cfqp8wfv5z.jpg,#dadadc,#d1d1d1,#d0d0d0
https://i.redd.it/7ofdqmp53bzy.jpg,#958a6e,#94896d,#93886c
https://i.redd.it/ftd3sx5ah13z.jpg,#fdffff,#b44b2c,#b54b31
https://i.redd.it/lsuw4p2ncyny.jpg,#ffffff,#a77842,#fffffe
http://i.imgur.com/enhDnTM.jpg,#ffffff,#000000,#c7e5ef
2019/04/01 22:23:18 totalTime:  5.644842085s
$ go run three-colors.go > tc.csv
2019/04/01 22:33:15 using too much memory. pausing. (Alloc = 133 MiB)
2019/04/01 22:33:16 using too much memory. pausing. (Alloc = 346 MiB)
2019/04/01 22:33:16 using too much memory. pausing. (Alloc = 166 MiB)
2019/04/01 22:33:17 using too much memory. pausing. (Alloc = 147 MiB)
2019/04/01 22:33:17 using too much memory. pausing. (Alloc = 130 MiB)
2019/04/01 22:33:18 using too much memory. pausing. (Alloc = 130 MiB)
2019/04/01 22:33:20 totalTime:  6.23480342s
$ ll
total 40
total 40
drwxr-xr-x   7 gregz  staff   224 Apr  1 22:28 .
drwxr-xr-x   4 gregz  staff   128 Apr  1 10:32 ..
drwxr-xr-x  13 gregz  staff   416 Apr  1 22:34 .git
-rw-r--r--   1 gregz  staff   192 Apr  1 10:30 .gitignore
-rw-r--r--   1 gregz  staff  3340 Apr  1 22:34 README.md
-rw-r--r--   1 gregz  staff  2305 Apr  1 22:33 tc.csv
-rw-r--r--   1 gregz  staff  6043 Apr  1 22:33 three-colors.go
```


