package utils

import "errors"

// 轮询算法
type RoundPollBalance struct {
	curIndex int
	rss      []string
}

/**
 * @Author: yang
 * @Description：添加服务
 * @Date: 2021/4/7 15:36
 */
func (r *RoundPollBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("params len 1 at least")
	}
	for _, v := range params {
		r.rss = append(r.rss, v)
	}

	return nil
}

/**
 * @Author: yang
 * @Description：轮询获取服务
 * @Date: 2021/4/7 15:36
 */
func (r *RoundPollBalance) Next() string {
	if len(r.rss) == 0 {
		return ""
	}
	lens := len(r.rss)
	if r.curIndex >= lens {
		r.curIndex = 0
	}
	curAdd := r.rss[r.curIndex]
	r.curIndex = (r.curIndex + 1) % lens
	return curAdd
}
