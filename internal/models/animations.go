package models

import (
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/nirabyte/todo/internal/config"
	"github.com/nirabyte/todo/internal/styles"
	"github.com/nirabyte/todo/internal/themes"
)

func renderCheckAnim(t Task, theme themes.Theme) string {
	elapsed := time.Since(t.AnimStart).Seconds()
	total := config.CheckAnimDuration.Seconds()
	progress := elapsed / total
	if progress > 1.0 {
		progress = 1.0
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	text := t.Title
	fallBack := lipgloss.NewStyle().Foreground(theme.Success).Render(text)

	switch t.AnimType {
	case AnimSparkle:
		chars := []string{"*", "+", "°", ".", "x", "o"}
		var sb strings.Builder
		for i := 0; i < len(text); i++ {
			if r.Float32() < 0.4 {
				char := chars[r.Intn(len(chars))]
				col := theme.Accent
				if r.Intn(2) == 0 {
					col = theme.Secondary
				}
				sb.WriteString(lipgloss.NewStyle().Foreground(col).Render(char))
			} else {
				sb.WriteString(lipgloss.NewStyle().Foreground(theme.Dim).Render(string(text[i])))
			}
		}
		return sb.String()
	case AnimMatrix:
		matrixChars := "H3LL0W0RLD$#@!%*&^"
		var sb strings.Builder
		for i := 0; i < len(text); i++ {
			char := string(matrixChars[r.Intn(len(matrixChars))])
			sb.WriteString(lipgloss.NewStyle().Foreground(theme.Success).Render(char))
		}
		return sb.String()
	case AnimWipeRight:
		idx := int(math.Floor(progress * float64(len(text))))
		var sb strings.Builder
		for i := 0; i < len(text); i++ {
			if i < idx {
				sb.WriteString(styles.StrikeStyle.Render(string(text[i])))
			}
			if i == idx {
				sb.WriteString(lipgloss.NewStyle().Background(theme.Secondary).Foreground(theme.Bg).Render(string(text[i])))
			}
			if i > idx {
				sb.WriteString(string(text[i]))
			}
		}
		return sb.String()
	case AnimWipeLeft:
		idx := len(text) - 1 - int(math.Floor(progress*float64(len(text))))
		var sb strings.Builder
		for i := 0; i < len(text); i++ {
			if i > idx {
				sb.WriteString(styles.StrikeStyle.Render(string(text[i])))
			}
			if i == idx {
				sb.WriteString(lipgloss.NewStyle().Background(theme.Accent).Foreground(theme.Bg).Render(string(text[i])))
			}
			if i < idx {
				sb.WriteString(string(text[i]))
			}
		}
		return sb.String()
	case AnimRainbow:
		colors := []lipgloss.Color{theme.Accent, theme.Secondary, theme.Success, theme.Warning, "#FF0000", "#00FF00", "#0000FF"}
		var sb strings.Builder
		for _, char := range text {
			c := colors[r.Intn(len(colors))]
			sb.WriteString(lipgloss.NewStyle().Foreground(c).Render(string(char)))
		}
		return sb.String()
	case AnimWave:
		colors := []lipgloss.Color{theme.Accent, theme.Secondary, theme.Success, theme.Fg}
		offset := int(elapsed * 30)
		var sb strings.Builder
		for i, char := range text {
			cIdx := (i + offset) % len(colors)
			sb.WriteString(lipgloss.NewStyle().Foreground(colors[cIdx]).Render(string(char)))
		}
		return sb.String()
	case AnimBinary:
		var sb strings.Builder
		for i := 0; i < len(text); i++ {
			bit := "0"
			if r.Intn(2) == 1 {
				bit = "1"
			}
			sb.WriteString(lipgloss.NewStyle().Foreground(theme.Success).Render(bit))
		}
		return sb.String()
	case AnimDissolve:
		var sb strings.Builder
		for i := 0; i < len(text); i++ {
			if r.Float64() < progress*1.5 {
				sb.WriteString(styles.StrikeStyle.Render(string(text[i])))
			} else {
				sb.WriteString(string(text[i]))
			}
		}
		return sb.String()
	case AnimFlip:
		var sb strings.Builder
		for _, char := range text {
			s := string(char)
			if r.Float32() < 0.3 {
				if strings.ToUpper(s) == s {
					s = strings.ToLower(s)
				} else {
					s = strings.ToUpper(s)
				}
				sb.WriteString(lipgloss.NewStyle().Foreground(theme.Secondary).Render(s))
			} else {
				sb.WriteString(s)
			}
		}
		return sb.String()
	case AnimPulse:
		var sb strings.Builder
		phase := math.Sin(elapsed * 40)
		col := theme.Fg
		if phase > 0 {
			col = theme.Accent
		}
		for _, char := range text {
			sb.WriteString(lipgloss.NewStyle().Foreground(col).Bold(phase > 0).Render(string(char)))
		}
		return sb.String()
	case AnimTypewriter:
		visibleChars := int(float64(len(text)) * progress)
		var sb strings.Builder
		for i := 0; i < len(text); i++ {
			if i <= visibleChars {
				sb.WriteString(lipgloss.NewStyle().Foreground(theme.Success).Render(string(text[i])))
			} else {
				sb.WriteString(" ")
			}
		}
		return sb.String()
	case AnimParticle:
		var sb strings.Builder
		for i := 0; i < len(text); i++ {
			if r.Float32() < 0.5 {
				sb.WriteString(lipgloss.NewStyle().Foreground(theme.Secondary).Render("."))
			} else {
				sb.WriteString(lipgloss.NewStyle().Foreground(theme.Accent).Render(string(text[i])))
			}
		}
		return sb.String()
	case AnimRedact:
		var sb strings.Builder
		chars := []string{"█", "▓", "▒", "░"}
		for i := 0; i < len(text); i++ {
			if r.Float32() < 0.5 {
				char := chars[r.Intn(len(chars))]
				sb.WriteString(lipgloss.NewStyle().Foreground(theme.Warning).Render(char))
			} else {
				sb.WriteString(lipgloss.NewStyle().Foreground(theme.Dim).Render(string(text[i])))
			}
		}
		return sb.String()
	case AnimChaos:
		symbols := "!@#$%^&*()_+"
		var sb strings.Builder
		for i := 0; i < len(text); i++ {
			if r.Float32() < 0.5 {
				s := string(symbols[r.Intn(len(symbols))])
				sb.WriteString(lipgloss.NewStyle().Foreground(theme.Secondary).Render(s))
			} else {
				sb.WriteString(string(text[i]))
			}
		}
		return sb.String()
	case AnimConverge:
		var sb strings.Builder
		mid := len(text) / 2
		fill := int(float64(mid) * progress)
		for i := 0; i < len(text); i++ {
			if i < fill || i >= len(text)-fill {
				sb.WriteString(styles.StrikeStyle.Render(string(text[i])))
			} else {
				sb.WriteString(lipgloss.NewStyle().Foreground(theme.Accent).Render(string(text[i])))
			}
		}
		return sb.String()
	case AnimBounce:
		var sb strings.Builder
		for i, char := range text {
			if r.Intn(2) == 0 {
				sb.WriteString(lipgloss.NewStyle().Foreground(theme.Accent).Render(string(char)))
			} else {
				sb.WriteString(lipgloss.NewStyle().Foreground(theme.Secondary).Render(string(char)))
			}
			if i%2 == int(elapsed*10)%2 {
				sb.WriteString("")
			}
		}
		return sb.String()
	case AnimSpin:
		spinners := []string{"-", "\\", "|", "/"}
		spinIdx := int(elapsed*20) % 4
		var sb strings.Builder
		for range text {
			sb.WriteString(lipgloss.NewStyle().Foreground(theme.Success).Render(spinners[spinIdx]))
		}
		return sb.String()
	case AnimZipper:
		var sb strings.Builder
		mid := len(text) / 2
		zipperPos := int(progress * float64(mid))
		for i := 0; i < len(text); i++ {
			distFromEdge := i
			if i >= mid {
				distFromEdge = len(text) - 1 - i
			}

			if distFromEdge < zipperPos {
				sb.WriteString(styles.StrikeStyle.Render(string(text[i])))
			} else {
				sb.WriteString(lipgloss.NewStyle().Background(theme.Accent).Foreground(theme.Bg).Render(string(text[i])))
			}
		}
		return sb.String()
	case AnimEraser:
		var sb strings.Builder
		for i := 0; i < len(text); i++ {
			if r.Float64() < progress {
				sb.WriteString(" ")
			} else {
				sb.WriteString(lipgloss.NewStyle().Foreground(theme.Dim).Render(string(text[i])))
			}
		}
		return sb.String()
	case AnimGlitch:
		var sb strings.Builder
		glitchChars := "¡¢£¤¥¦§¨©ª«¬®¯°±²³´µ¶·¸¹º»¼½¾¿"
		for i := 0; i < len(text); i++ {
			if r.Float32() < 0.3 {
				char := string(glitchChars[r.Intn(len(glitchChars))])
				sb.WriteString(lipgloss.NewStyle().Foreground(theme.Warning).Background(theme.Dim).Render(char))
			} else {
				sb.WriteString(string(text[i]))
			}
		}
		return sb.String()
	case AnimMoons:
		phases := []string{"◐", "◓", "◑", "◒"}
		var sb strings.Builder
		for i := 0; i < len(text); i++ {
			if r.Float32() < 0.3 {
				sb.WriteString(lipgloss.NewStyle().Foreground(theme.Secondary).Render(phases[r.Intn(len(phases))]))
			} else {
				sb.WriteString(lipgloss.NewStyle().Foreground(theme.Dim).Render(string(text[i])))
			}
		}
		return sb.String()
	case AnimBraille:
		var sb strings.Builder
		for i := 0; i < len(text); i++ {
			if r.Float32() < 0.4 {
				sb.WriteString(lipgloss.NewStyle().Foreground(theme.Accent).Render(string(rune(0x2800 + r.Intn(255)))))
			} else {
				sb.WriteString(string(text[i]))
			}
		}
		return sb.String()
	case AnimHex:
		var sb strings.Builder
		hexChars := "0123456789ABCDEF"
		for i := 0; i < len(text); i++ {
			sb.WriteString(lipgloss.NewStyle().Foreground(theme.Success).Render(string(hexChars[r.Intn(len(hexChars))])))
		}
		return sb.String()
	case AnimReverse:
		var sb strings.Builder
		for i := len(text) - 1; i >= 0; i-- {
			sb.WriteString(lipgloss.NewStyle().Foreground(theme.Warning).Render(string(text[i])))
		}
		return sb.String()
	case AnimCaseFlip:
		var sb strings.Builder
		for _, char := range text {
			s := string(char)
			if r.Intn(2) == 0 {
				s = strings.ToUpper(s)
			} else {
				s = strings.ToLower(s)
			}
			sb.WriteString(lipgloss.NewStyle().Foreground(theme.Accent).Render(s))
		}
		return sb.String()
	case AnimWide:
		var sb strings.Builder
		for _, char := range text {
			sb.WriteString(lipgloss.NewStyle().Foreground(theme.Secondary).Render(string(char) + " "))
		}
		return sb.String()
	case AnimTraffic:
		colors := []lipgloss.Color{theme.Warning, "#FFFF00", theme.Success}
		cIdx := int(elapsed*10) % 3
		return lipgloss.NewStyle().Foreground(colors[cIdx]).Render(text)
	case AnimCenterStrike:
		var sb strings.Builder
		mid := len(text) / 2
		strikeWidth := int(progress * float64(mid))
		for i := 0; i < len(text); i++ {
			dist := int(math.Abs(float64(i - mid)))
			if dist < strikeWidth {
				sb.WriteString(styles.StrikeStyle.Render(string(text[i])))
			} else {
				sb.WriteString(string(text[i]))
			}
		}
		return sb.String()
	case AnimLoading:
		var sb strings.Builder
		fill := int(progress * float64(len(text)))
		for i := 0; i < len(text); i++ {
			if i < fill {
				sb.WriteString(lipgloss.NewStyle().Foreground(theme.Success).Render("█"))
			} else {
				sb.WriteString(lipgloss.NewStyle().Foreground(theme.Dim).Render("▒"))
			}
		}
		return sb.String()
	case AnimSlider:
		var sb strings.Builder
		for i := 0; i < len(text); i++ {
			if r.Float32() < 0.3 {
				sb.WriteString(lipgloss.NewStyle().Foreground(theme.Accent).Render("^"))
			} else {
				sb.WriteString(string(text[i]))
			}
		}
		return sb.String()
	}

	return fallBack
}

func renderDeleteAnim(text string, theme themes.Theme) string {
	var sb strings.Builder
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len(text); i++ {
		bit := "0"
		if r.Intn(2) == 1 {
			bit = "1"
		}
		sb.WriteString(lipgloss.NewStyle().Foreground(theme.Warning).Bold(true).Render(bit))
	}
	return sb.String()
}

