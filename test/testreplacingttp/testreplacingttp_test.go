package testreplacingttp_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"placeholder_elasticsearch/datamodels"
)

var _ = Describe("Testreplacingttp", func() {
	Context("Тест 1. Проверка замены старых значений TtpMessageEs объекта, новыми значениями, если они отличаются", func() {
		oldStruct := datamodels.TtpsMessageEs{
			Ttp: map[string][]datamodels.TtpMessage{
				"T1587.001": {
					{
						OccurDate:            "2024-01-22T11:53:00+03:00",
						UnderliningCreatedAt: "2024-01-22T11:54:03+03:00",
						UnderliningId:        "~1106653328",
						UnderliningCreatedBy: "zsiem@rcm.lcl",
						PatternId:            "T1587.001",
						Tactic:               "resource-development",
						ExtraData: datamodels.ExtraDataTtpMessage{
							Pattern: datamodels.PatternExtraData{
								RemoteSupport:        false,
								Revoked:              false,
								UnderliningCreatedAt: "2021-10-04T04:21:35+03:00",
								UnderliningCreatedBy: "admin@thehive.local",
								UnderliningId:        "~4890664",
								UnderliningType:      "Pattern",
								Description:          "Adversaries may develop malware and malware components that can be used during targeting. Building malicious software can include the development of payloads, droppers, post-compromise tools, backdoors (including backdoored images), packers, C2 protocols, and the creation of infected removable media. Adversaries may develop malware to support their operations, creating a means for maintaining control of remote machines, evading defenses, and executing post-compromise behaviors.(Citation: Mandiant APT1)(Citation: Kaspersky Sofacy)(Citation: ActiveMalwareEnergy)(Citation: FBI Flash FIN7 USB)As with legitimate development efforts, different skill sets may be required for developing malware. The skills needed may be located in-house, or may need to be contracted out. Use of a contractor may be considered an extension of that adversary's malware development capabilities, provided the adversary plays a role in shaping requirements and maintains a degree of exclusivity to the malware.Some aspects of malware development, such as C2 protocol development, may require adversaries to obtain additional infrastructure. For example, malware developed that will communicate with Twitter for C2, may require use of [Web Services](https://attack.mitre.org/techniques/T1583/006).(Citation: FireEye APT29)",
								Detection:            "Much of this activity will take place outside the visibility of the target organization, making detection of this behavior difficult. Detection efforts may be focused on post-compromise phases of the adversary lifecycle.",
								Name:                 "Malware",
								PatternId:            "T1587.001",
								PatternType:          "attack-pattern",
								URL:                  "https://attack.mitre.org/techniques/T1587/001",
								Version:              "1.1",
								Platforms: []string{
									"PRE",
									"Linux",
									"macOS",
									"Windows",
								},
								PermissionsRequired: []string{"User"},
								DataSources:         []string(nil),
								Tactics: []string{
									"resource-development",
									"command-and-control",
								},
							},
							PatternParent: datamodels.PatternExtraData{
								RemoteSupport:        false,
								Revoked:              false,
								UnderliningCreatedAt: "2021-10-04T04:19:44+03:00",
								UnderliningCreatedBy: "admin@thehive.local",
								UnderliningId:        "~4485160",
								UnderliningType:      "Pattern",
								Description:          "Adversaries may build capabilities that can be used during targeting. Rather than purchasing, freely downloading, or stealing capabilities, adversaries may develop their own capabilities in-house. This is the process of identifying development requirements and building solutions such as malware, exploits, and self-signed certificates. Adversaries may develop capabilities to support their operations throughout numerous phases of the adversary lifecycle.(Citation: Mandiant APT1)(Citation: Kaspersky Sofacy)(Citation: Bitdefender StrongPity June 2020)(Citation: Talos Promethium June 2020)As with legitimate development efforts, different skill sets may be required for developing capabilities. The skills needed may be located in-house, or may need to be contracted out. Use of a contractor may be considered an extension of that adversary's development capabilities, provided the adversary plays a role in shaping requirements and maintains a degree of exclusivity to the capability.",
								Detection:            "Much of this activity will take place outside the visibility of the target organization, making detection of this behavior difficult. Detection efforts may be focused on related stages of the adversary lifecycle, such as during Defense Evasion or Command and Control.",
								Name:                 "Develop Capabilities",
								PatternId:            "T1587",
								PatternType:          "attack-pattern",
								URL:                  "https://attack.mitre.org/techniques/T1587",
								Version:              "1.0",
								Platforms:            []string{"PRE"},
								PermissionsRequired:  []string(nil),
								DataSources:          []string(nil),
								Tactics:              []string{"resource-development"},
							},
						},
					},
				},
				"T1132": {
					{
						OccurDate:            "2024-01-25T21:12:01+03:00",
						UnderliningCreatedAt: "2024-01-25T03:00:51+03:00",
						UnderliningId:        "~1106665616",
						UnderliningCreatedBy: "zsiem@rcm.lcl",
						PatternId:            "T1132",
						Tactic:               "command-and-control",
						ExtraData: datamodels.ExtraDataTtpMessage{
							Pattern: datamodels.PatternExtraData{
								RemoteSupport:        false,
								Revoked:              false,
								UnderliningCreatedAt: "2021-09-04T10:09:41+03:00",
								UnderliningCreatedBy: "admin@thehive.local",
								UnderliningId:        "~45539408",
								UnderliningType:      "Pattern",
								Description:          "Adversaries may encode data to make the content of command and control traffic more difficult to detect. Command and control (C2) information can be encoded using a standard data encoding system. Use of data encoding may adhere to existing protocol specifications and includes use of ASCII, Unicode, Base64, MIME, or other binary-to-text and character encoding systems.(Citation: Wikipedia Binary-to-text Encoding) (Citation: Wikipedia Character Encoding) Some data encoding systems may also result in data compression, such as gzip.",
								Detection:            "Analyze network data for uncommon data flows (e.g., a client sending significantly more data than it receives from a server). Processes utilizing the network that do not normally have network communication or have never been seen before are suspicious. Analyze packet contents to detect communications that do not follow the expected protocol behavior for the port that is being used. (Citation: University of Birmingham C2)",
								Name:                 "Data Encoding",
								PatternId:            "T1132",
								PatternType:          "attack-pattern",
								URL:                  "https://attack.mitre.org/techniques/T1132",
								Version:              "1.1",
								Platforms:            []string(nil),
								PermissionsRequired:  []string(nil),
								DataSources:          []string{"Network Traffic: Network Traffic Content"},
								Tactics:              []string(nil),
							},
							PatternParent: datamodels.PatternExtraData{},
						},
					},
				},
			},
		}

		newStruct := datamodels.TtpsMessageEs{
			Ttp: map[string][]datamodels.TtpMessage{
				"T1587.001": {
					{
						//OccurDate:            "2024-01-22T11:53:00+03:00",
						UnderliningCreatedAt: "2024-01-02T00:15:03+03:00", //замена
						UnderliningId:        "~1106653328",
						UnderliningCreatedBy: "zsiem@rcm.lcl",
						PatternId:            "T1587.001",
						//Tactic:               "resource-development",
						ExtraData: datamodels.ExtraDataTtpMessage{
							Pattern: datamodels.PatternExtraData{
								RemoteSupport:        true, //замена
								Revoked:              false,
								UnderliningCreatedAt: "2021-10-04T04:21:35+03:00",
								UnderliningCreatedBy: "admin@thehive.local.local", //замена
								UnderliningId:        "~4890664",
								UnderliningType:      "Pattern element", //замена
								//Description:          "Adversaries may develop malware and malware components that can be used during targeting. Building malicious software can include the development of payloads, droppers, post-compromise tools, backdoors (including backdoored images), packers, C2 protocols, and the creation of infected removable media. Adversaries may develop malware to support their operations, creating a means for maintaining control of remote machines, evading defenses, and executing post-compromise behaviors.(Citation: Mandiant APT1)(Citation: Kaspersky Sofacy)(Citation: ActiveMalwareEnergy)(Citation: FBI Flash FIN7 USB)As with legitimate development efforts, different skill sets may be required for developing malware. The skills needed may be located in-house, or may need to be contracted out. Use of a contractor may be considered an extension of that adversary's malware development capabilities, provided the adversary plays a role in shaping requirements and maintains a degree of exclusivity to the malware.Some aspects of malware development, such as C2 protocol development, may require adversaries to obtain additional infrastructure. For example, malware developed that will communicate with Twitter for C2, may require use of [Web Services](https://attack.mitre.org/techniques/T1583/006).(Citation: FireEye APT29)",
								Detection: "Much of this activity will take place outside the visibility of the target organization, making detection of this behavior difficult. Detection efforts may be focused on post-compromise phases of the adversary lifecycle.",
								Name:      "Malware example", //замена
								PatternId: "T1587.001",
								//PatternType:          "attack-pattern",
								URL:     "https://attack.mitre.org/techniques/T1587/001",
								Version: "1.2", //замена
								//добавление
								Platforms: []string{
									"PRE",
									"Experement OS",
									"Linux",
									"macOS",
									"Windows",
								},
								//добавление
								PermissionsRequired: []string{
									"User",
									"root",
								},
								DataSources: []string(nil),
								Tactics: []string{
									"resource-development",
									"command-and-control",
								},
							},
							PatternParent: datamodels.PatternExtraData{
								RemoteSupport:        false,
								Revoked:              false,
								UnderliningCreatedAt: "2021-10-04T04:19:44+03:00",
								UnderliningCreatedBy: "admin@thehive.local",
								UnderliningId:        "~4485160",
								UnderliningType:      "Pattern",
								Description:          "Adversaries may build capabilities that can be used during targeting. Rather than purchasing, freely downloading, or stealing capabilities, adversaries may develop their own capabilities in-house. This is the process of identifying development requirements and building solutions such as malware, exploits, and self-signed certificates. Adversaries may develop capabilities to support their operations throughout numerous phases of the adversary lifecycle.(Citation: Mandiant APT1)(Citation: Kaspersky Sofacy)(Citation: Bitdefender StrongPity June 2020)(Citation: Talos Promethium June 2020)As with legitimate development efforts, different skill sets may be required for developing capabilities. The skills needed may be located in-house, or may need to be contracted out. Use of a contractor may be considered an extension of that adversary's development capabilities, provided the adversary plays a role in shaping requirements and maintains a degree of exclusivity to the capability.",
								Detection:            "Much of this activity will take place outside the visibility of the target organization, making detection of this behavior difficult. Detection efforts may be focused on related stages of the adversary lifecycle, such as during Defense Evasion or Command and Control.",
								Name:                 "Develop Capabilities",
								PatternId:            "T1587",
								PatternType:          "attack-pattern",
								URL:                  "https://attack.mitre.org/techniques/T1587",
								Version:              "1.0",
								Platforms:            []string{"PRE"},
								PermissionsRequired:  []string(nil),
								DataSources:          []string(nil),
								Tactics:              []string{"resource-development"},
							},
						},
					},
					//добавление
					{
						OccurDate:            "2024-03-03T11:33:00+03:00",
						UnderliningCreatedAt: "2024-03-04T16:30:15+03:00",
						UnderliningId:        "~2204766328",
						UnderliningCreatedBy: "zsiem@rcm.lcl",
						PatternId:            "T1587.001",
						Tactic:               "resource-development",
						ExtraData: datamodels.ExtraDataTtpMessage{
							Pattern: datamodels.PatternExtraData{
								RemoteSupport:        false,
								Revoked:              false,
								UnderliningCreatedAt: "2021-10-04T04:21:35+03:00",
								UnderliningCreatedBy: "admin@thehive.local",
								UnderliningId:        "~5079288",
								UnderliningType:      "Pattern",
								Description:          "Adversaries may buy",
								Detection:            "Much of this activity will take place",
								Name:                 "Malware",
								PatternId:            "T1587.001",
								PatternType:          "attack-pattern",
								URL:                  "https://attack.mitre.org/techniques/T1588/001",
								Version:              "1.0",
								Platforms:            []string{"PRE"},
								PermissionsRequired:  []string(nil),
								DataSources:          []string(nil),
								Tactics:              []string{"resource-development"},
							},
							PatternParent: datamodels.PatternExtraData{
								RemoteSupport:        false,
								Revoked:              false,
								UnderliningCreatedAt: "2021-10-04T04:20:08+03:00",
								UnderliningCreatedBy: "admin@thehive.local",
								UnderliningId:        "~45645904",
								UnderliningType:      "Pattern",
								Description:          "Adversaries may buy and/or steal capabilities",
								Detection:            "Much of this activity will take place outside",
								Name:                 "Obtain Capabilities",
								PatternId:            "T1587",
								PatternType:          "attack-pattern",
								URL:                  "https://attack.mitre.org/techniques/T1587",
								Version:              "1.0",
								Platforms: []string{
									"PRE",
									"MSOffice",
									"Windows OS",
								},
								PermissionsRequired: []string{"Admin"},
								DataSources:         []string(nil),
								Tactics: []string{
									"resource-development",
									"application-development",
								},
							},
						},
					},
				},
				"T1132": {
					{
						OccurDate:            "2024-01-15T01:32:01+03:00", //замена
						UnderliningCreatedAt: "2024-01-21T03:41:51+03:00", //замена
						UnderliningId:        "~1106665616",
						UnderliningCreatedBy: "zsiem_example@rcm.lcl", //замена
						PatternId:            "T1132",
						Tactic:               "command-and-control",
						ExtraData: datamodels.ExtraDataTtpMessage{
							Pattern: datamodels.PatternExtraData{
								RemoteSupport:        false,
								Revoked:              true,                                //замена
								UnderliningCreatedAt: "2021-09-14T14:19:41+03:00",         //замена
								UnderliningCreatedBy: "admin_example_email@thehive.local", //замена
								UnderliningId:        "~45539408",
								UnderliningType:      "Pattern example", //замена
								Description:          "Adversaries may encode data to make the content of command and control traffic more difficult to detect. Command and control (C2) information can be encoded using a standard data encoding system. Use of data encoding may adhere to existing protocol specifications and includes use of ASCII, Unicode, Base64, MIME, or other binary-to-text and character encoding systems.(Citation: Wikipedia Binary-to-text Encoding) (Citation: Wikipedia Character Encoding) Some data encoding systems may also result in data compression, such as gzip.",
								Detection:            "Analyze network data for uncommon data flows (e.g., a client sending significantly more data than it receives from a server). Processes utilizing the network that do not normally have network communication or have never been seen before are suspicious. Analyze packet contents to detect communications that do not follow the expected protocol behavior for the port that is being used. (Citation: University of Birmingham C2)",
								Name:                 "Data Encoding Element", //замена
								PatternId:            "T1132",
								PatternType:          "attack-pattern",
								URL:                  "https://attack.mitre.org/techniques/T1132_update_link", //замена
								Version:              "1.1",
								//добавление
								Platforms: []string{"MacOS", "LinuxOS"},
								//добавление
								PermissionsRequired: []string{"User", "SimpleUser", "Admin"},
								DataSources:         []string{"Network Traffic: Network Traffic Content"},
								//добавление
								Tactics: []string{"impact"},
							},
							//добавление
							PatternParent: datamodels.PatternExtraData{
								RemoteSupport:        false,
								Revoked:              false,
								UnderliningCreatedAt: "2021-10-04T04:19:47+03:00",
								UnderliningCreatedBy: "admin@thehive.local",
								UnderliningId:        "~45437080",
								UnderliningType:      "Pattern",
								Description:          "Adversaries may perform Endpoint Denial of Service (DoS)",
								Detection:            "Detection of Endpoint DoS can sometimes be achieved before the effect is sufficient to cause",
								Name:                 "Endpoint Denial of Service",
								PatternId:            "T1499",
								PatternType:          "attack-pattern",
								URL:                  "https://attack.mitre.org/techniques/T1499",
								Version:              "1.1",
								Platforms: []string{
									"Windows",
									"Azure AD",
									"Office 365",
									"SaaS",
									"IaaS",
									"Linux",
									"macOS",
									"Google Workspace",
									"Containers",
								},
								PermissionsRequired: []string(nil),
								DataSources: []string{
									"Sensor Health: Host Status",
									"Application Log: Application Log Content",
									"Network Traffic: Network Traffic Content",
									"Network Traffic: Network Traffic Flow",
								},
								Tactics: []string{"impact"},
							},
						},
					},
				},
			},
		}

		It("Ряд полей в TtpMessageEs должны быть успешно заменены", func() {
			num := oldStruct.ReplacingOldValues(newStruct)

			//кол-во замененных полей
			Expect(num).Should(Equal(42))

			fmt.Println("---=== VERIFED TtpMessageEs ===---")
			fmt.Println(oldStruct.ToStringBeautiful(0))

			T1587001, ok := oldStruct.GetKeyTtp("T1587.001")
			Expect(ok).Should(BeTrue())
			Expect(len(T1587001)).Should(Equal(2))
			Expect(T1587001[0].GetUnderliningCreatedAt()).Should(Equal("2024-01-02T00:15:03+03:00"))
			Expect(T1587001[0].GetPattern().GetRemoteSupport()).Should(BeTrue())
			Expect(T1587001[0].GetPattern().GetUnderliningCreatedBy()).Should(Equal("admin@thehive.local.local"))
			Expect(T1587001[0].GetPattern().GetUnderliningType()).Should(Equal("Pattern element"))
			Expect(T1587001[0].GetPattern().GetName()).Should(Equal("Malware example"))
			Expect(len(T1587001[0].GetPattern().GetPlatforms())).Should(Equal(5))
			Expect(T1587001[0].GetPatternParent().GetUnderliningId()).Should(Equal("~4485160"))

			Expect(T1587001[1].GetUnderliningId()).Should(Equal("~2204766328"))

			T1132, ok := oldStruct.GetKeyTtp("T1132")
			Expect(ok).Should(BeTrue())
			Expect(T1132[0].GetOccurDate()).Should(Equal("2024-01-15T01:32:01+03:00"))
			Expect(T1132[0].GetUnderliningCreatedAt()).Should(Equal("2024-01-21T03:41:51+03:00"))
			Expect(T1132[0].GetUnderliningCreatedBy()).Should(Equal("zsiem_example@rcm.lcl"))
			Expect(T1132[0].GetPattern().UnderliningCreatedBy).Should(Equal("admin_example_email@thehive.local"))

			Expect(true).Should(BeTrue())
		})
	})
})
