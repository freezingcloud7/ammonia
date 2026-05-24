package main

import (
	"fmt"
	"time"
)

// CageModule represents a detachable, lightweight modular storage unit.
type CageModule struct {
	MountedLocation string // "Back", "Shoulder", "Thigh"
	MethaneAmount   int
	AmmoniaBolts    int
}

// SpaceSuit manages the resource metrics and life-support subsystems for the space warrior.
type SpaceSuit struct {
	WarriorName          string
	SuitBattery          float64
	BodyTemperature      float64
	GlassBolts           int     // Cryogenic Saturn-class ammonia glass bolts
	FireBolts            int     // Ignite-ready thermal bio-mass bolts
	MethaneGas           int     // Methane propulsion gas (%)
	SolidWasteBuffer     int     // Solid bio-mass reserve (%)
	VinylStock           int     // Anti-stick nano-vinyl membrane inventory (units)
	MaxTankCapacity      int     
	SaturnAtmosphereSync float64 // Atmospheric chemical synchronization rate (%)
	BackupCages          []CageModule
}

// 1. VerifyAmmoniaPurity runs high-resolution telemetry to ensure 100% match with naturally occurring Saturnian Ammonia Glass.
func (suit *SpaceSuit) VerifyAmmoniaPurity() {
	fmt.Println("\n[Telemetry] 🔬 Launching elemental analysis on internal Ammonia reserves...")
	time.Sleep(500 * time.Millisecond)
	
	suit.SaturnAtmosphereSync = 100.0
	fmt.Printf("[Status] 🪐 Synchronization optimal: 100%% matches Saturnian Ammonia Glass composition.\n")
	fmt.Println("[Security] Structural integrity verified. Kinetic armor-piercing matrix initialized.")
}

// 2. ProcessSolidWaste utilizes premium anti-stick nano-vinyl to safely encapsulate solid waste into thermal projectiles.
func (suit *SpaceSuit) ProcessSolidWaste() {
	if suit.SolidWasteBuffer < 50 || suit.VinylStock <= 0 {
		fmt.Println("\n[System] ⚠️ Fabrication suspended: Insufficient solid bio-mass or Anti-Stick Vinyl stock.")
		return
	}

	fmt.Println("\n[System] 💩 Solid bio-mass detected. Initializing Clean-Packaging Protocol...")
	time.Sleep(500 * time.Millisecond)

	suit.SolidWasteBuffer -= 50
	suit.VinylStock--
	suit.FireBolts += 2 

	fmt.Println("[Fabrication] Space-grade Anti-Stick Nano-Vinyl applied. Zero-residue containment secured.")
	fmt.Printf("[Inventory] Encapsulated Thermal Bio-Fuel Projectiles +2 generated. (Current Fire Bolts: %d)\n", suit.FireBolts)
}

// 3. RechargeBioBattery diverts solid bio-mass into the combustion fuel cell to emergency-charge the core power grid.
func (suit *SpaceSuit) RechargeBioBattery() {
	if suit.SolidWasteBuffer < 40 {
		fmt.Println("\n[⚡ Power Grid] Recharge aborted: Insufficient bio-mass for combustion.")
		return
	}

	fmt.Println("\n[System] 🔌 Main battery critical. Routing raw organic assets to the combustion fuel cell...")
	time.Sleep(500 * time.Millisecond)

	suit.SolidWasteBuffer -= 40
	suit.SuitBattery = minValue(100.0, suit.SuitBattery+35.5)

	fmt.Println("[Energy] Microbial combustion loop stabilized. Power grid synchronization complete.")
	fmt.Printf("[Status] Suit main power core successfully restored to: %.1f%%\n", suit.SuitBattery)
}

// 4. ProcessFlatulence processes gas into methane propulsion and routes overflows into modular backup cages.
func (suit *SpaceSuit) ProcessFlatulence() {
	fmt.Println("\n[System] 💨 Gaseous bio-mass detected. Sorting into methane and ammonia lines...")
	suit.MethaneGas += 80
	suit.GlassBolts += 4

	if suit.MethaneGas > suit.MaxTankCapacity {
		overflowMethane := suit.MethaneGas - suit.MaxTankCapacity
		suit.MethaneGas = suit.MaxTankCapacity
		
		locations := []string{"Left Thigh", "Right Thigh", "Left Shoulder", "Right Shoulder"}
		currentCageCount := len(suit.BackupCages)
		
		var targetLoc string
		if currentCageCount < len(locations) {
			targetLoc = locations[currentCageCount]
		} else {
			targetLoc = "Upper Back Pack"
		}

		newCage := CageModule{
			MountedLocation: targetLoc,
			MethaneAmount:   overflowMethane,
			AmmoniaBolts:    2,
		}
		suit.BackupCages = append(suit.BackupCages, newCage)
		fmt.Printf("[📦 Logistics] Primary tank capacity exceeded. Auxiliary lightweight cage mounted on [%s]\n", targetLoc)
	}
}

// 5. FireWeapon manages the external dual-element pistol pipeline, handling both ICE and FIRE firing vectors.
package main

import (
	"fmt"
	"time"
)

// CageModule represents a detachable, lightweight modular storage unit.
type CageModule struct {
	MountedLocation string // "Back", "Shoulder", "Thigh"
	MethaneAmount   int
	AmmoniaBolts    int
}

// SpaceSuit manages the resource metrics and life-support subsystems for the space warrior.
type SpaceSuit struct {
	WarriorName          string
	SuitBattery          float64
	BodyTemperature      float64
	GlassBolts           int     // Cryogenic Saturn-class ammonia glass bolts
	FireBolts            int     // Ignite-ready thermal bio-mass bolts
	MethaneGas           int     // Methane propulsion gas (%)
	SolidWasteBuffer     int     // Solid bio-mass reserve (%)
	VinylStock           int     // Anti-stick nano-vinyl membrane inventory (units)
	MaxTankCapacity      int
	SaturnAtmosphereSync float64 // Atmospheric chemical synchronization rate (%)
	BackupCages          []CageModule
}

// 1. VerifyAmmoniaPurity runs high-resolution telemetry to ensure 100% match with naturally occurring Saturnian Ammonia Glass.
func (suit *SpaceSuit) VerifyAmmoniaPurity() {
	fmt.Println("\n[Telemetry] 🔬 Launching elemental analysis on internal Ammonia reserves...")
	time.Sleep(500 * time.Millisecond)
	suit.SaturnAtmosphereSync = 100.0
	fmt.Printf("[Status] 🪐 Synchronization optimal: 100%% matches Saturnian Ammonia Glass composition.\n")
	fmt.Println("[Security] Structural integrity verified. Kinetic armor-piercing matrix initialized.")
}

// 2. ProcessSolidWaste utilizes premium anti-stick nano-vinyl to safely encapsulate solid waste into thermal projectiles.
func (suit *SpaceSuit) ProcessSolidWaste() {
	if suit.SolidWasteBuffer < 50 || suit.VinylStock <= 0 {
		fmt.Println("\n[System] ⚠️ Fabrication suspended: Insufficient solid bio-mass or Anti-Stick Vinyl stock.")
		return
	}
	fmt.Println("\n[System] 💩 Solid bio-mass detected. Initializing Clean-Packaging Protocol...")
	time.Sleep(500 * time.Millisecond)
	suit.SolidWasteBuffer -= 50
	suit.VinylStock--
	suit.FireBolts += 2
	fmt.Println("[Fabrication] Space-grade Anti-Stick Nano-Vinyl applied. Zero-residue containment secured.")
	fmt.Printf("[Inventory] Encapsulated Thermal Bio-Fuel Projectiles +2 generated. (Current Fire Bolts: %d)\n", suit.FireBolts)
}

// 3. RechargeBioBattery diverts solid bio-mass into the combustion fuel cell to emergency-charge the core power grid.
func (suit *SpaceSuit) RechargeBioBattery() {
	if suit.SolidWasteBuffer < 40 {
		fmt.Println("\n[⚡ Power Grid] Recharge aborted: Insufficient bio-mass for combustion.")
		return
	}
	fmt.Println("\n[System] 🔌 Main battery critical. Routing raw organic assets to the combustion fuel cell...")
	time.Sleep(500 * time.Millisecond)
	suit.SolidWasteBuffer -= 40
	suit.SuitBattery = minValue(100.0, suit.SuitBattery+35.5)
	fmt.Println("[Energy] Microbial combustion loop stabilized. Power grid synchronization complete.")
	fmt.Printf("[Status] Suit main power core successfully restored to: %.1f%%\n", suit.SuitBattery)
}

// 4. ProcessFlatulence processes gas into methane propulsion and routes overflows into modular backup cages.
func (suit *SpaceSuit) ProcessFlatulence() {
	fmt.Println("\n[System] 💨 Gaseous bio-mass detected. Sorting into methane and ammonia lines...")
	suit.MethaneGas += 80
	suit.GlassBolts += 4
	if suit.MethaneGas > suit.MaxTankCapacity {
		overflowMethane := suit.MethaneGas - suit.MaxTankCapacity
		suit.MethaneGas = suit.MaxTankCapacity
		locations := []string{"Left Thigh", "Right Thigh", "Left Shoulder", "Right Shoulder"}
		currentCageCount := len(suit.BackupCages)
		var targetLoc string
		if currentCageCount < len(locations) {
			targetLoc = locations[currentCageCount]
		} else {
			targetLoc = "Upper Back Pack"
		}
		newCage := CageModule{
			MountedLocation: targetLoc,
			MethaneAmount:   overflowMethane,
			AmmoniaBolts:    2,
		}
		suit.BackupCages = append(suit.BackupCages, newCage)
		fmt.Printf("[📦 Logistics] Primary tank capacity exceeded. Auxiliary lightweight cage mounted on [%s]\n", targetLoc)
	}
}

// 5. FireWeapon manages the external dual-element pistol pipeline, handling both ICE and FIRE firing vectors.
func (suit *SpaceSuit) FireWeapon(mode string) {
	switch mode {
	case "ICE":
		if suit.GlassBolts <= 0 {
			fmt.Println("\n[Weapon Error] Discharge failed: Out of Ammonia Glass Bolts!")
			return
		}
		suit.GlassBolts--
		fmt.Println("\n[Discharge] 🥶 SHATTER! Launching Saturnian Cryogenic Ammonia Glass Bolt!")
		fmt.Println("[Impact] Kinetic energy delivered. Target neutralized with extreme cryogenic flash freeze.")
	case "FIRE":
		if suit.FireBolts <= 0 {
			fmt.Println("\n[Weapon Error] Discharge failed: Out of Bio Fire Bolts!")
			return
		}
		suit.FireBolts--
		fmt.Println("\n[Discharge] 🔥 IGNITION! Launching Clean-Packaged Solid Bio Fire Bolt!")
		fmt.Println("[Survival] Thermal ignition successful. Projectile thermal yield sufficient for planetary campfire deploy.")
	}
	fmt.Printf("[Weapon Status] Active Munitions - Ice Bolts: %d | Fire Bolts: %d\n", suit.GlassBolts, suit.FireBolts)
}

func minValue(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Initialize tactical squad unit profile
	warrior := SpaceSuit{
		WarriorName:      "Titan-Vanguard",
		SuitBattery:      15.0,
		BodyTemperature:  36.5,
		GlassBolts:       2,
		FireBolts:        0,
		MethaneGas:       30,
		SolidWasteBuffer: 100,
		VinylStock:       5,
		MaxTankCapacity:  100,
	}

	fmt.Println("======================================================================")
	fmt.Println("🪐 Saturnian Ammonia Glass Bolt & Bio-Loop Space Suit OS v2.0.0")
	fmt.Println("======================================================================")

	// Execute core operating system boot routine
	warrior.VerifyAmmoniaPurity()

	// Execute autonomous life-support recycling loops
	warrior.ProcessFlatulence()
	warrior.ProcessSolidWaste()
	warrior.RechargeBioBattery()

	// Execute weapons system simulation test
	warrior.FireWeapon("ICE")
	warrior.FireWeapon("FIRE")

	// ======================================================================
	// 🌌 v2.0.0 HUD & DIAGNOSTICS ENGINES LINKED INSIDE MAIN
	// ======================================================================
	hudDisplay := VirtualHUD{}
	warrior.RunDiagnostics(&hudDisplay)

	fmt.Println("\n======================================================================")
	fmt.Println("[Execution Success] All universal eco-laws and planetary physics applied.")
	fmt.Println("======================================================================")
}

// ======================================================================
// 🌌 NEXT ARCHITECTURE: v2.0.0 HIGH-TECH RECIPE (ADD BELOW MAIN)
// ======================================================================

// VirtualHUD manages the zero-hardware AR screen and real-time medical diagnostics.
type VirtualHUD struct {
	IsVisible          bool
	ActiveWaterControl float64 // Advanced moisture and hydration control level (%)
	WarriorHealthScore int     // Real-time medical health index (0-100)
}

// RunDiagnostics analyzes the molecules filtered by the MEFF system to perform an instantaneous medical checkup.
func (suit *SpaceSuit) RunDiagnostics(hud *VirtualHUD) {
	fmt.Println("\n[HUD OS] 🕶️ Projecting Hardware-Free Virtual AR Screen directly onto the helmet visor...")
	time.Sleep(500 * time.Millisecond)
	hud.IsVisible = true

	fmt.Println("[MEFF Link] 🔬 Analyzing bio-molecular components from the fractionation filter...")
	time.Sleep(500 * time.Millisecond)

	// Simulated moisture control and medical evaluation
	hud.ActiveWaterControl = 98.7
	hud.WarriorHealthScore = 95

	fmt.Println("\n================== 🩺 REAL-TIME MEDICAL DIAGNOSTICS ==================")
	fmt.Printf(" [VITALS] Warrior Health Score: %d/100 (Status: EXCELLENT)\n", hud.WarriorHealthScore)
	fmt.Printf(" [HYDRO] Automated Moisture & Hydration Control: %.1f%% Sync\n", hud.ActiveWaterControl)
	fmt.Println(" [DIAGNOSIS] Urea/Ammonia levels stable. No signs of cosmic radiation exposure.")
	fmt.Println(" [BIO-SECURITY] Human resources are 100% secured and optimized for combat.")
	fmt.Println("=======================================================================")
}