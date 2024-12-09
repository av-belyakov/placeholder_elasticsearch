package testhandlerobservables

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchSensorIdFromDescription(t *testing.T) {
	var str string = "**Задача переданная из смежной системы: Заслон-Пост-Модерн**В формате ГЦМ: **`TSK-8MSK-6-ZPM-241201-1956748`** ID: `1956748`[http://siem.cloud.gcm/tasks/card/TSK-8MSK-6-ZPM-241201-1956748](http://siem.cloud.gcm/tasks/card/TSK-8MSK-6-ZPM-241201-1956748)Автор задачи: **`Security Event Manager`**Тип: **`snort_alert`****Причина по которой создана задача**Название: `Редко встречающиеся признаки ВПО, LFI  attempts с 103.109.101.6`Описание: `## Данная задача создана автоматически Время начала: 2024-12-01 00:15:30 Время окончания: 2024-12-01 00:15:30 Продолжительность воздействий: 0:00:00`Отработало на СОА: - **`8030030`**   ОБ Портал, Установлен: Москва,Москва, IP адрес: 10.20.0.30**Полное описание события IDS:**- Время начала: **`01.12.2024 00:15:30`**- Время окончания: **`01.12.2024 00:15:30`**- **IP из домашней подсети**1. **`213.24.76.23`**- **IP из внешней подсети**1. **`103.109.101.6`****Сигнатуры на которых отработал анализатор сетевого трафика:**1. РП: **`53994872`**, Сообщение: Trojan-Downloader.Agent.HTTP.C&C, Добавлена: 29.05.2024 12:22:51**Ссылка на arkime:**  - [http://arkime.cloud.gcm/2024_11_30_21_22_30_894042____1733001745_2024_12_01____00_22_25_936014.pcap](http://anisimova.cloud.gcm:8005/sessions?expression=file%20%3D%3D%20%2Fopt%2Farkime%2Fraw%2F2024_11_30_21_22_30_894042____1733001745_2024_12_01____00_22_25_936014.pcap&date=-1)"

	rexSensorId := regexp.MustCompile(`СОА:\s-\s\*\*\x60(\d+)\x60\*\*`)
	tmp := rexSensorId.FindStringSubmatch(str)

	fmt.Println("TMP:", tmp)

	assert.Equal(t, "8030030", tmp[1])
}
