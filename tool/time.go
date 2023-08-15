package tool

import (
	"time"
)

type TimeTool struct {
	timeStr string
	format TimeFormat
	loc *time.Location
	tm *time.Time
}

func (t *TimeTool) NewTimeTool(timeStr string, format TimeFormat) *TimeTool  {
	if timeStr!= "" {
		t.timeStr = timeStr
	} else {
		if format != "" {
			t.timeStr = time.Now().Format(string(format))
		} else {
			t.timeStr = time.Now().Format(string(Y_M_D_H_I_S))
		}
	}

	t.format = format
	return t
}

func (t *TimeTool) LoadLocation(loc TimeLocation) *TimeTool  {

	t.loc, _ = time.LoadLocation(string(loc))

	return t
}

func (t *TimeTool) AddDay(day int) *TimeTool  {

	if t.format == "" {
		t.format = Y_M_D_H_I_S
	}

	if t.tm == nil {
		if t.loc != nil {
			if t.timeStr != "" {
				tm, _ := time.ParseInLocation(string(t.format), t.timeStr, t.loc)
				tm = tm.AddDate(0,0, day)
				t.tm = &tm
			} else {
				tm := time.Now().AddDate(0,0, day)
				t.tm = &tm
			}


		} else {
			if t.timeStr != "" {
				tm, _ := time.Parse(string(t.format), t.timeStr)
				tm = tm.AddDate(0,0, day)
				t.tm = &tm
			} else {
				tm := time.Now().AddDate(0,0, day)
				t.tm = &tm
			}

		}
	} else {
		tm := t.tm.AddDate(0,0, day)
		t.tm = &tm
	}
	


	return t
}

func (t *TimeTool) AddMonth(month int) *TimeTool {

	if t.format == "" {
		t.format = Y_M_D_H_I_S
	}

	if t.tm == nil {
		if t.loc != nil {
			if t.timeStr != "" {
				tm, _ := time.ParseInLocation(string(t.format), t.timeStr, t.loc)
				tm = tm.AddDate(0,month, 0)
				t.tm = &tm
			} else {
				tm := time.Now().AddDate(0,month, 0)
				t.tm = &tm
			}


		} else {
			if t.timeStr != "" {
				tm, _ := time.Parse(string(t.format), t.timeStr)
				tm = tm.AddDate(0,month, 0)
				t.tm = &tm
			} else {
				tm := time.Now().AddDate(0,month, 0)
				t.tm = &tm
			}

		}
	} else {
		tm := t.tm.AddDate(0,month, 0)
		t.tm = &tm
	}

	return t
}

func (t *TimeTool) AddYear(year int) *TimeTool  {

	if t.format == "" {
		t.format = Y_M_D_H_I_S
	}

	if t.tm == nil {
		if t.loc != nil {
			if t.timeStr != "" {
				tm, _ := time.ParseInLocation(string(t.format), t.timeStr, t.loc)
				tm = tm.AddDate(year,0, 0)
				t.tm = &tm
			} else {
				tm := time.Now().In(t.loc).AddDate(year,0, 0)
				t.tm = &tm
			}


		} else {
			if t.timeStr != "" {
				tm, _ := time.Parse(string(t.format), t.timeStr)
				tm = tm.AddDate(year,0, 0)
				t.tm = &tm
			} else {
				tm := time.Now().AddDate(year,0, 0)
				t.tm = &tm
			}

		}
	} else {
		tm := t.tm.AddDate(year,0, 0)
		t.tm = &tm
	}

	return t
}

func (t *TimeTool) AddSecond(second int) *TimeTool  {

	if t.format == "" {
		t.format = Y_M_D_H_I_S
	}

	if t.tm == nil {
		if t.loc != nil {
			if t.timeStr != "" {
				tm, _ := time.ParseInLocation(string(t.format), t.timeStr, t.loc)
				tm = tm.Add(time.Duration(second))
				t.tm = &tm
			} else {
				tm := time.Now().In(t.loc).Add(time.Duration(second))
				t.tm = &tm
			}

		} else {
			if t.timeStr != "" {
				tm, _ := time.Parse(string(t.format), t.timeStr)
				tm = tm.Add(time.Duration(second))
				t.tm = &tm
			} else {
				tm := time.Now().Add(time.Duration(second))
				t.tm = &tm
			}

		}
	} else {
		tm := t.tm.Add(time.Duration(second))
		t.tm = &tm
	}

	return t
}

func (t *TimeTool) GetTime() (tm time.Time)  {

	if t.format == "" {
		t.format = Y_M_D_H_I_S
	}

	if t.loc != nil {
		if t.tm == nil {
			if t.timeStr != "" {

				tm, _  = time.ParseInLocation(string(t.format), t.timeStr, t.loc)
			} else {
				tm = time.Now().In(t.loc)
			}

		} else {
			tm = t.tm.In(t.loc)
		}

	} else {
		if t.tm == nil {
			if t.timeStr != "" {
				tm, _  = time.Parse( string(t.format), t.timeStr)
			} else {
				tm = time.Now()
			}

		} else {
			tm = *t.tm
		}
	}

	return tm
}

func (t *TimeTool) GetString() (tm string) {

	if t.format == "" {
		t.format = Y_M_D_H_I_S
	}

	if t.loc != nil {
		if t.tm == nil {
			if t.timeStr != "" {
				tm  = t.timeStr
			} else {
				tm = time.Now().In(t.loc).Format(string(t.format))
			}
		} else {
			tm = t.tm.In(t.loc).Format(string(t.format))
		}

	} else {

		if t.tm == nil {
			if t.timeStr != "" {
				tm  = t.timeStr
			} else {
				tm = time.Now().Format(string(t.format))
			}
		} else {
			tm = t.tm.Format(string(t.format))
		}
	}

	return tm
}
