package workspace

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Responce(s string) string {
	var result string
	// все предложение в нижний регистр
	s = strings.ToLower(s)
	tkm := strings.Contains(s, "тк")
	zavis := strings.Contains(s, "завис")
	visit := strings.Contains(s, "висит")
	pk := strings.Contains(s, "комп")
	igr := strings.Contains(s, "игровой")
	neR := strings.Contains(s, "не реагирует")
	zayavka := strings.Contains(s, "заявка")
	program := strings.Contains(s, "софт")

	if program {
		result = `Имеется 
		IVMS-4200,
		VMS tools, 
		KMS-Auto,
		KonBoot для 7,
		Драйвер на локалку
		
		`
	}

	if tkm && zavis || visit && tkm || igr && zavis || igr && tkm {
		result = "Посмотрите интернет кабеля "
	}
	if pk && zavis || pk && visit {
		result = "Перезагрузите комп"
	}
	if neR && tkm {
		result = "Проверьте кабель питания, если так же то заеду скоро посмотрю"
	}
	if zayavka {
		result = "Хорошо заеду"
	}

	return result
}

func Command(s string) string {
	var result string
	switch s {
	case "/a":
		result = `Если не можете нажать кнопку включения то посмотрите ВИДЕО👇
			Если не помогает, то проверьте кабеля питания сзади монитора или под столом
			
		`
	case "/t":
		result = `Если по всему залу не работает или зависает трансляция на теликах и игровых
		то нужно перезагрузить РЕСИВЕР👆 на 10 минут минимум
		чтобы перезагрузить сзади слева есть толстый кабель питания его нужно отключить
		
		`
	case "/nt":
		result = `
			если в таком 👆 положении то нужно подтолкнуть кабеля интернет белые либо синие и перезагрузить
			Если все равно не будет грузиться то нужно проверить включен системный блок(ПРОЦЕССОР) в серверной
			Иначе напишите мне Wa.me/77004010700
			
		`
	case "/p":
		result = `Нужно подтолкнуть интернет кабеля на ТК
		`
	case "/s":
		result = `
			Если есть надпись в верхнем правом углу "No touch device"
			то переподключите usb👆  на ТК
		`
	case "/nm":
		result = `
			Если у вас он моргает как СВЕТОФОР то нажмите на мониторе  кнопку imput и выберите "Auto" либо DVI
			Если он не включается то проверьте кабеля питания черный толстый сзади или под столом
		`
	case "/vm":
		result = `
			если монитор не включается Проверьте кабель питания черный толстый 
			если монитор работает то проверьте черный плоский кабель, подтолкните его и перезагрузите ТК
		`
	case "/f1":
		result = `Подключите клавиатуру к ТК, войдите в BIOS и выставьте следующие настройки. Чтобы войти в БИОС, нажмите F1, esc или del.
		1. Заходим во вкладку Advanced
		2. Далее открываем OnBoard Device Configuration
		3. В пункте OnBoard Lan Boot Rom выбираем пункт Enabled.
		4. Далее заходим во вкладку Boot.
		5. В пункте Boot Option #1 выбираем пункт Network.
		6. Проверьте, что в пункте Boot Mode select был параметр Legacy. Измените на него по необходимости с UEFI.
		7. Нажимаем кнопку F4, сохраняем данные.`

	}

	return result
}

func Video(chatID int64, name string, vName string) tgbotapi.VideoConfig {
	video := tgbotapi.NewVideo(chatID, tgbotapi.FilePath("Video/"+vName+".mp4"))
	video.Thumb = tgbotapi.FilePath("Video/" + vName + ".mp4")
	return video
}
