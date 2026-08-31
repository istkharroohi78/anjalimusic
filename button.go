package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

// Button ke styles define kiye gaye hain
type ButtonStyle string

const (
	Success   ButtonStyle = "success"
	Primary   ButtonStyle = "primary"
	Danger    ButtonStyle = "danger"
	Secondary ButtonStyle = "secondary"
	Info      ButtonStyle = "info"
)

// StyledButton function inline buttons generate karta hai
func StyledButton(text string, callbackData string, url string, style ButtonStyle) tb.Btn {
	// Agar URL diya hai toh URL button banega
	if url != "" {
		return tb.Btn{
			Text: text,
			URL:  url,
		}
	}

	// Agar Callback Data diya hai toh Data button banega
	if callbackData != "" {
		return tb.Btn{
			Text: text,
			Data: callbackData,
		}
	}

	// Default fallback
	return tb.Btn{
		Text: text,
		Data: "none",
	}
}
