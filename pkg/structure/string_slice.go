package structure

type StringSlice []string

func (s *StringSlice) FilterNil() *StringSlice {
    res := make([]string, 0)
    for _, v := range *s {
        if v != "" {
            res = append(res, v)
        }
    }
    if len(res) > 0 {
        *s = res
    }
    return s
}

func (s StringSlice) FilterNilAndCreate() []string {
    res := make([]string, 0)
    for _, v := range s {
        if v != "" {
            res = append(res, v)
        }
    }
    return res
}

func (s StringSlice) Is(str string) bool {
    for _, v := range s {
        if v == str {
            return true
        }
    }
    return false
}

func (s StringSlice) Output() []string {
    return s
}

func (s StringSlice) Length() int {
    return len(s)
}
