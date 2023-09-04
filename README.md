# grrrrrrr

## Demonstration
📷 [Image Source](https://www.pinterest.com/pin/345510602680940816/)

<div style="display: grid; grid-template-columns: 1fr;">
  <img src="https://i.imgur.com/UX8Xtyo.jpg" alt="source image" style="width: 96%;">
</div>

<div style="display: grid; grid-template-columns: 1fr 1fr;">
  <img src="https://i.imgur.com/FKM1Au0.png" alt="Image 3" style="width: 96%;">
</div>

<div style="display: grid; grid-template-columns: 1fr 1fr;">
  <img src="https://i.imgur.com/BkuCdR9.png" alt="Image 1" style="width: 48%;">
  <img src="https://i.imgur.com/ku441go.png" alt="Image 2" style="width: 48%;">
</div>

<div style="display: grid; grid-template-columns: 1fr 1fr;">
  <img src="https://i.imgur.com/oLkJH4k.png" alt="Image 3" style="width: 48%;">
  <img src="https://i.imgur.com/JzWOces.png" alt="Image 4" style="width: 48%;">
</div>

<div style="display: grid; grid-template-columns: 1fr 1fr;">
  <img src="https://i.imgur.com/ro6Au1j.png" alt="Image 3" style="width: 48%;">
  <img src="https://i.imgur.com/tS4GP9b.png" alt="Image 4" style="width: 48%;">
</div>

## CLI Flags
- src `string` <br>
*Source image to be converted<br>
.png / .jpeg*<br>
(default [sample image](assets/images/fallen-angels-1995.jpeg))
- dest `string`<br>
*Destination directory<br>
(default current directory)*
- method `string`<br>
    - conversion methods:
      - average 
      - luma
      - desaturate
      - decomposeMin
      - decomposeMax
      - singleChannel
      - shades ( might wanna use this with a blurring effect to smoothen out the image e.g. [gauswuchs](https://github.com/lulzshadowwalker/gauswuchs)
<br> [see implementation comments for description](pkg/grrrrrrr/grrrrrrr.go)
- colorChannel `string`<br>
*determines which color channel to use for --method="singleChannel"<br>
(default red)*
- shadeCount `int`<br>
*determines the number of shades when using the [shades] conversion method<br>
(default 8)*
- colorChannel `string`<br>
*determines which color channel to use for --method="singleChannel"<br>
(default red)*
- shadeCount `int`<br>
*determines the number of shades when using the [shades] conversion method<br>
(default 8)*
- help <br>
*Lists all commands*
