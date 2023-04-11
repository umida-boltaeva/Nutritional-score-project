package main

import "fmt"

func main()  {
	ns := GetNutritionalScore(NutritionalData{
		Energy: EnergyFromKcal(),
		Sugar: SugarGram(),
		SaturatedFattyAcids: SaturatedFattyAcids(),
		Sodium: SodiumMilligram(),
		Fruits: FruitsPercent(),
		Protein: ProteinGram(),
	}, Food)
	fmt.Printf("Nutritional Score: %d\n", ns.Value)
}