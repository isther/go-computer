## Half Adder

和数：<img src="https://latex.codecogs.com/svg.image?A\oplus&space;B" title="https://latex.codecogs.com/svg.image?A\oplus B" />

进位：<img src="https://latex.codecogs.com/svg.image?AB" title="https://latex.codecogs.com/svg.image?AB" />

![](../../image/half_adder.png)

|  A  |  B  | S0  | C0  |
| :-: | :-: | :-: | :-: |
|  0  |  0  |  0  |  0  |
|  1  |  0  |  1  |  0  |
|  0  |  1  |  1  |  0  |
|  1  |  1  |  0  |  1  |

## Full Adder

和数：<img src="https://latex.codecogs.com/svg.image?A_i\oplus&space;B_i\oplus&space;C_{i-1}"/>

进位：<img src="https://latex.codecogs.com/svg.image?A_iB_i&plus;C_{i-1}(A_i&plus;B_i)"/>

进位也可用一个异或门来代替或门对其中两个输入信号进行求和<img src="https://latex.codecogs.com/svg.image?A_iB_i&plus;C_{i-1}(A_i\oplus&space;B_i)"/>

![](../../image/full_adder.png)

| $C_i$ |  A  |  B  | $S_i$ | $C_i$ |
| :---: | :-: | :-: | :---: | :---: |
|   0   |  0  |  0  |   0   |   0   |
|   1   |  0  |  0  |   1   |   0   |
|   0   |  1  |  0  |   1   |   0   |
|   1   |  1  |  0  |   0   |   1   |
|   0   |  0  |  1  |   1   |   0   |
|   1   |  0  |  1  |   0   |   1   |
|   0   |  1  |  1  |   0   |   1   |
|   1   |  1  |  1  |   1   |   1   |

## Four Bit Adder

![](../../image/four_bit_adder.png)
