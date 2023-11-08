package main

const (
	HELP = "Данный бот имеет следующие возможности:\n\n" +
		"1. Определять язык введённого текста. \nШаблон входных данных:\n\n" + "`/определить | <текст>`\n\n" +
		"2. Переводить введённый текст (можно вводить текст, не зная, на каком языке он введён; бот сам скажет, на каком языке введён ваш текст). " +
		"\nШаблон входных данных:\n\n" + "/перевести | <код желаемого языка> | <текст>\n" +
		"\n>>Языки, которые на данный момент может распознать бот при переводе: русский, английский, немецкий, французский, испанский, португальский, арабский, китайский, японский и хинди."
)