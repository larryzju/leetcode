func intToRoman(num int) string {
        res := []byte{}
        fn := func(i, v, x byte) func(n int) []byte {
                return func(n int) []byte {
                        switch n {
                        case 1: return []byte{i}
                        case 2: return []byte{i,i}
                        case 3: return []byte{i,i,i}
                        case 4: return []byte{i,v}
                        case 5: return []byte{v}
                        case 6: return []byte{v,i}
                        case 7: return []byte{v,i,i}
                        case 8: return []byte{v,i,i,i}
                        case 9: return []byte{i,x}
                        default: return []byte{}
                        }
                }
        }

        b1 := fn('I', 'V', 'X')
        b2 := fn('X', 'L', 'C')
        b3 := fn('C', 'D', 'M')

        for num >= 1000 {
                res = append(res, 'M')
                num -= 1000
        }

        res = append(res, b3(num / 100 % 10)...)
        res = append(res, b2(num / 10 % 10)...)
        res = append(res, b1(num % 10)...)
        return string(res)
}
