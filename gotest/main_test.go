//go:generate abigen --pkg gotest -sol ../contracts/Crowdsale.sol -out ./contract.go

package gotest

import (
	"math/big"
	"testing"
	"time"
)

const stageDuration = 600

func TestDeploy(t *testing.T) {
	t.Run("pure deploy is successful", func(t *testing.T) {
		sim, user := StartEthereumSimulator(1)
		if sim == nil || user == nil {
			t.Fatal("Simulator should have started")
		}
		addr := DeployContract(user[0], sim, big.NewInt(time.Now().Unix()), big.NewInt(stageDuration), big.NewInt(stageDuration), big.NewInt(stageDuration), user[0].Opts.From, user[0].Opts.From)
		session := GetCrowdsaleSession(user[0], sim, addr)
		if session == nil {
			t.Fatal("Should successfuly get contract session")
		}
	})

	t.Run("State is valid after deployment", func(t *testing.T) {
		sim, user := StartEthereumSimulator(2)
		if sim == nil || user == nil {
			t.Fatal("Simulator should have started")
		}
		addr := DeployContract(user[0], sim, big.NewInt(time.Now().Unix()+2), big.NewInt(stageDuration), big.NewInt(stageDuration), big.NewInt(stageDuration), user[0].Opts.From, user[0].Opts.From)
		session := GetCrowdsaleSession(user[0], sim, addr)
		if session == nil {
			t.Fatal("Should successfuly get contract session")
		}
		tokenSession := GetTokenSession(user[0], sim, addr)
		if tokenSession == nil {
			t.Fatal("Token session should not be nil")
		}

		t.Run("State is valid", func(t *testing.T) {

			if decimals, _ := tokenSession.Decimals(); decimals.Int64() != 8 {
				t.Fatal("Token decimals should be 8")
			}

			if wallet, _ := session.Wallet(); wallet.Hash() != user[0].Opts.From.Hash() {
				t.Fatal("Wallets should be equal")
			}

			if fiatSymbol, _ := session.FiatSymbol(); fiatSymbol != "EUR" {
				t.Fatal("Fiat symbol should be EUR")
			}

			if fiatDecimals, _ := session.FiatDecimals(); fiatDecimals.Int64() != 2 {
				t.Fatal("Fiat decimals should be 2")
			}

			if state, _ := session.CrowdsaleState(); state != 0 {
				t.Fatal("Right after the deploy state should be 0 (NOT_STARTED)")
			}

			if currentHardCap, _ := session.CurrentHardCap(); currentHardCap.Int64() != 0 {
				t.Fatal("Current hard cap should be 0 because crowdsale is not started yet")
			}
		})

		t.Run("Crowdsale is fully initialized after prepareCrowdsale()", func(t *testing.T) {

			tx, err := session.PrepareCrowdsale()
			if err != nil {
				t.Fatal("Failed to send TX prepareCrowdsale")
			}
			receipt := GetReceipt(tx, sim)
			if receipt.Status != 1 {
				t.Fatal("TX should not fail")
			}

			if privPrice, _ := session.PrivatePriceInFiatFracture(); privPrice.Int64() != 25 {
				t.Fatal("Price for Private Sale should be 25 (0.25 EUR)")
			}

			if prePrice, _ := session.PreIcoPriceInFiatFracture(); prePrice.Int64() != 50 {
				t.Fatal("Price for Pre-ICO should be 50 (0.5 EUR)")
			}

			if icoPrice, _ := session.IcoPriceInFiatFracture(); icoPrice.Int64() != 100 {
				t.Fatal("Price for Private Sale should be 100 (1 EUR)")
			}

			tokenUnit := big.NewInt(10)
			tokenUnit.Exp(tokenUnit, big.NewInt(8), nil) // 1E8
			expectedGeneralHardCap := big.NewInt(160000000)
			expectedGeneralHardCap.Mul(expectedGeneralHardCap, tokenUnit)
			if generalHardCap, _ := session.GeneralHardCap(); generalHardCap.Cmp(expectedGeneralHardCap) != 0 {
				t.Fatal("General hard cap should be 160 mln tokens")
			}

			expectedPrivateHardCap := big.NewInt(10000000)
			expectedPrivateHardCap.Mul(expectedPrivateHardCap, tokenUnit)
			if privHardCap, _ := session.PrivateSaleHardCap(); privHardCap.Cmp(expectedPrivateHardCap) != 0 {
				t.Fatal("Private Sale hard cap should be 10 mln tokens")
			}

			expectedPreIcoHardCap := big.NewInt(30000000)
			expectedPreIcoHardCap.Mul(expectedPreIcoHardCap, tokenUnit)
			if preHardCap, _ := session.PreIcoHardCap(); preHardCap.Cmp(expectedPreIcoHardCap) != 0 {
				t.Fatal("Pre-ICO hard cap should be 30 mln tokens")
			}

			expectedIcoHardCap := big.NewInt(60000000)
			expectedIcoHardCap.Mul(expectedIcoHardCap, tokenUnit)
			if icoHardCap, _ := session.IcoHardCap(); icoHardCap.Cmp(expectedIcoHardCap) != 0 {
				t.Fatal("ICO hard cap should be 60 mln tokens")
			}
		})
	})
}

func TestOwnership(t *testing.T) {
	sim, user := StartEthereumSimulator(2)
	if sim == nil || user == nil {
		t.Fatal("Simulator should have started")
	}
	addr := DeployContract(user[0], sim, big.NewInt(time.Now().Unix()+2), big.NewInt(stageDuration), big.NewInt(stageDuration), big.NewInt(stageDuration), user[0].Opts.From, user[0].Opts.From)
	//owner session
	session1 := GetCrowdsaleSession(user[0], sim, addr)
	if session1 == nil {
		t.Fatal("Should successfuly get contract session")
	}
	tokenSession1 := GetTokenSession(user[0], sim, addr)
	if tokenSession1 == nil {
		t.Fatal("Token session should not be nil")
	}

	//non-owner session
	session2 := GetCrowdsaleSession(user[1], sim, addr)
	if session2 == nil {
		t.Fatal("Should successfuly get contract session")
	}
	tokenSession2 := GetTokenSession(user[1], sim, addr)
	if tokenSession2 == nil {
		t.Fatal("Token session should not be nil")
	}

	t.Run("not owner", func(t *testing.T) {
		if owner, _ := session2.Owner(); owner.Hash() == user[1].Opts.From.Hash() {
			t.Fatal("User 1 should not be owner of the crowdsale")
		}

		tx, err := session2.PrepareCrowdsale()
		if err != nil {
			t.Fatal("Failed to send TX prepareCrowdsale")
		}
		receipt := GetReceipt(tx, sim)
		if receipt.Status != 0 {
			t.Fatal("TX should fail because of onlyOwner modifier")
		}

		tx, err = session2.AddPrivateParticipant(user[1].Opts.From)
		if err != nil {
			t.Fatal("Failed to send TX AddPrivateParticipant")
		}
		receipt = GetReceipt(tx, sim)
		if receipt.Status != 0 {
			t.Fatal("TX should fail because of onlyOwner modifier")
		}

		tx, err = session2.SetWallet(user[1].Opts.From)
		if err != nil {
			t.Fatal("Failed to send TX SetWallet")
		}
		receipt = GetReceipt(tx, sim)
		if receipt.Status != 0 {
			t.Fatal("TX should fail because of onlyOwner modifier")
		}

		tx, err = session2.SetNewStartDate(big.NewInt(123456789))
		if err != nil {
			t.Fatal("Failed to send TX SetNewStartDate")
		}
		receipt = GetReceipt(tx, sim)
		if receipt.Status != 0 {
			t.Fatal("TX should fail because of onlyOwner modifier")
		}

		tx, err = session2.ManualReserve(user[1].Opts.From, big.NewInt(1000))
		if err != nil {
			t.Fatal("Failed to send TX ManualReserve")
		}
		receipt = GetReceipt(tx, sim)
		if receipt.Status != 0 {
			t.Fatal("TX should fail because of onlyOwner modifier")
		}

		tx, err = session2.Withdraw()
		if err != nil {
			t.Fatal("Failed to send TX Withdraw")
		}
		receipt = GetReceipt(tx, sim)
		if receipt.Status != 0 {
			t.Fatal("TX should fail because of onlyOwner modifier")
		}
	})

	t.Run("owner", func(t *testing.T) {
		if owner, _ := session1.Owner(); owner.Hash() == user[1].Opts.From.Hash() {
			t.Fatal("User 1 should not be owner of the crowdsale")
		}

		tx, err := session1.PrepareCrowdsale()
		if err != nil {
			t.Fatal("Failed to send TX prepareCrowdsale")
		}
		receipt := GetReceipt(tx, sim)
		if receipt.Status != 1 {
			t.Fatal("TX should not fail")
		}

		tx, err = session1.AddPrivateParticipant(user[1].Opts.From)
		if err != nil {
			t.Fatal("Failed to send TX AddPrivateParticipant")
		}
		receipt = GetReceipt(tx, sim)
		if receipt.Status != 1 {
			t.Fatal("TX should not fail")
		}

		tx, err = session1.SetWallet(user[1].Opts.From)
		if err != nil {
			t.Fatal("Failed to send TX SetWallet")
		}
		receipt = GetReceipt(tx, sim)
		if receipt.Status != 1 {
			t.Fatal("TX should not fail")
		}

		tx, err = session1.SetNewStartDate(big.NewInt(123456789))
		if err != nil {
			t.Fatal("Failed to send TX SetNewStartDate")
		}
		receipt = GetReceipt(tx, sim)
		if receipt.Status != 1 {
			t.Fatal("TX should not fail")
		}
	})
}

func TestCrowdsale(t *testing.T) {
	sim, user := StartEthereumSimulator(2)
	if sim == nil || user == nil {
		t.Fatal("Simulator should have started")
	}
	t.Run("Cannot buy before starting date", func(t *testing.T) {
		t.Run("Without adding to private sale participants", func(t *testing.T) {
			addr := DeployContract(user[0], sim, big.NewInt(time.Now().Unix()+2), big.NewInt(stageDuration), big.NewInt(stageDuration), big.NewInt(stageDuration), user[0].Opts.From, user[0].Opts.From)
			//owner session
			session := GetCrowdsaleSession(user[0], sim, addr)
			if session == nil {
				t.Fatal("Should successfully get contract session")
			}
			tokenSession := GetTokenSession(user[0], sim, addr)
			if tokenSession == nil {
				t.Fatal("Token session should not be nil")
			}

			//non-owner session
			session2 := GetCrowdsaleSession(user[1], sim, addr)
			if session2 == nil {
				t.Fatal("Should successfully get contract session")
			}
			tokenSession2 := GetTokenSession(user[1], sim, addr)
			if tokenSession2 == nil {
				t.Fatal("Token session should not be nil")
			}

			tx, err := session.PrepareCrowdsale()
			if err != nil {
				t.Fatal("Failed to send TX prepareCrowdsale")
			}
			receipt := GetReceipt(tx, sim)
			if receipt.Status != 1 {
				t.Fatal("TX should not fail")
			}

			tx = sendTransaction(user[1], sim, addr, "1000000000000000000")
			if tx == nil {
				t.Fatal("TX should have been created successfuly")
			}
			receipt = GetReceipt(tx, sim)
			if receipt.Status != 0 {
				t.Fatal("TX should have failed")
			}
		})

		t.Run("With adding to private sale participants", func(t *testing.T) {
			addr := DeployContract(user[0], sim, big.NewInt(time.Now().Unix()+2), big.NewInt(stageDuration), big.NewInt(stageDuration), big.NewInt(stageDuration), user[0].Opts.From, user[0].Opts.From)
			//owner session
			session := GetCrowdsaleSession(user[0], sim, addr)
			if session == nil {
				t.Fatal("Should successfuly get contract session")
			}
			tokenSession := GetTokenSession(user[0], sim, addr)
			if tokenSession == nil {
				t.Fatal("Token session should not be nil")
			}

			tx, err := session.PrepareCrowdsale()
			if err != nil {
				t.Fatal("Failed to send TX prepareCrowdsale")
			}
			receipt := GetReceipt(tx, sim)
			if receipt.Status != 1 {
				t.Fatal("TX should not fail")
			}

			//non-owner session
			session2 := GetCrowdsaleSession(user[1], sim, addr)
			if session2 == nil {
				t.Fatal("Should successfuly get contract session")
			}
			tokenSession2 := GetTokenSession(user[1], sim, addr)
			if tokenSession2 == nil {
				t.Fatal("Token session should not be nil")
			}

			tx, err = session.AddPrivateParticipant(user[1].Opts.From)
			if err != nil {
				t.Fatal("TX should have been successfully created")
			}
			receipt = GetReceipt(tx, sim)
			if receipt.Status != 1 {
				t.Fatal("Adding private participant should have been successful")
			}

			tx = sendTransaction(user[1], sim, addr, "1000000000000000000")
			if tx == nil {
				t.Fatal("Transaction should have been successfully created")
			}
			receipt = GetReceipt(tx, sim)
			if receipt.Status != 0 {
				t.Fatal("TX should still fail, because crowdsale hasn't started yet")
			}
		})
	})

	t.Run("Buy after start date", func(t *testing.T) {
		t.Run("With fiat price set", func(t *testing.T) {
			t.Run("Buy without being in the participants list", func(t *testing.T) {
				addr := DeployContract(user[0], sim, big.NewInt(time.Now().Unix()+2), big.NewInt(stageDuration), big.NewInt(stageDuration), big.NewInt(stageDuration), user[0].Opts.From, user[0].Opts.From)
				//owner session
				session := GetCrowdsaleSession(user[0], sim, addr)
				if session == nil {
					t.Fatal("Should successfuly get contract session")
				}
				tokenSession := GetTokenSession(user[0], sim, addr)
				if tokenSession == nil {
					t.Fatal("Token session should not be nil")
				}

				tx, err := session.PrepareCrowdsale()
				if err != nil {
					t.Fatal("Failed to send TX prepareCrowdsale")
				}
				receipt := GetReceipt(tx, sim)
				if receipt.Status != 1 {
					t.Fatal("TX should not fail")
				}

				tx, err = session.UpdateWeiInFiat(big.NewInt(1000000000000000)) // 0.001 ETH in 0.01 EUR
				if err != nil {
					t.Fatal("Transaction should have been successfully created")
				}
				receipt = GetReceipt(tx, sim)
				if receipt.Status != 1 {
					t.Fatal("TX should have been successful")
				}

				time.Sleep(3 * time.Second)

				tx = sendTransaction(user[1], sim, addr, "1000000000000000000")
				if tx == nil {
					t.Fatal("Transaction should have been successfully created")
				}
				receipt = GetReceipt(tx, sim)
				if receipt.Status != 0 {
					t.Fatal("TX should have failed, because sender is not in the private participants list")
				}

				balance, err := tokenSession.BalanceOf(user[1].Opts.From)
				if err != nil {
					t.Fatal("Transaction should have been successfully created")
				}
				if balance.Int64() != 0 {
					t.Fatal("Balance should be 0")
				}
			})

			t.Run("Buy if address is in participants list", func(t *testing.T) {
				addr := DeployContract(user[0], sim, big.NewInt(time.Now().Unix()+2), big.NewInt(stageDuration), big.NewInt(stageDuration), big.NewInt(stageDuration), user[0].Opts.From, user[0].Opts.From)
				//owner session
				session := GetCrowdsaleSession(user[0], sim, addr)
				if session == nil {
					t.Fatal("Should successfuly get contract session")
				}
				tokenSession := GetTokenSession(user[0], sim, addr)
				if tokenSession == nil {
					t.Fatal("Token session should not be nil")
				}

				tx, err := session.PrepareCrowdsale()
				if err != nil {
					t.Fatal("Failed to send TX prepareCrowdsale")
				}
				receipt := GetReceipt(tx, sim)
				if receipt.Status != 1 {
					t.Fatal("TX should not fail")
				}

				tx, err = session.UpdateWeiInFiat(big.NewInt(1000000000000000)) // 0.001 ETH in 0.01 EUR
				if err != nil {
					t.Fatal("Transaction should have been successfully created")
				}
				receipt = GetReceipt(tx, sim)
				if receipt.Status != 1 {
					t.Fatal("TX should have been successful")
				}

				tx, err = session.AddPrivateParticipant(user[1].Opts.From)
				if err != nil {
					t.Fatal("TX should have been successfully created")
				}
				receipt = GetReceipt(tx, sim)
				if receipt.Status != 1 {
					t.Fatal("TX should be successful")
				}

				time.Sleep(3 * time.Second)

				tx = sendTransaction(user[1], sim, addr, "1000000000000000000")
				if tx == nil {
					t.Fatal("Transaction should have been successfully created")
				}
				receipt = GetReceipt(tx, sim)
				if receipt.Status != 1 {
					t.Fatal("TX should be successful")
				}

				balance, err := tokenSession.BalanceOf(user[1].Opts.From)
				if err != nil {
					t.Fatal("Transaction should have been successfully created")
				}
				if balance.Int64() != 40 {
					t.Fatal("Balance should be 40, buy was successful with rate 0.001 ETH = 0.01 EUR and 1 ETH transfer")
				}
			})
		})

		t.Run("With fiat price not set", func(t *testing.T) {
			t.Run("Buy without being in the participants list", func(t *testing.T) {
				addr := DeployContract(user[0], sim, big.NewInt(time.Now().Unix()+2), big.NewInt(stageDuration), big.NewInt(stageDuration), big.NewInt(stageDuration), user[0].Opts.From, user[0].Opts.From)
				//owner session
				session := GetCrowdsaleSession(user[0], sim, addr)
				if session == nil {
					t.Fatal("Should successfuly get contract session")
				}
				tokenSession := GetTokenSession(user[0], sim, addr)
				if tokenSession == nil {
					t.Fatal("Token session should not be nil")
				}

				time.Sleep(3 * time.Second)

				tx := sendTransaction(user[1], sim, addr, "1000000000000000000")
				if tx == nil {
					t.Fatal("Transaction should have been successfully created")
				}
				receipt := GetReceipt(tx, sim)
				if receipt.Status != 0 {
					t.Fatal("TX should have failed, because sender is not in the private participants list")
				}

				balance, err := tokenSession.BalanceOf(user[1].Opts.From)
				if err != nil {
					t.Fatal("Transaction should have been successfully created")
				}
				if balance.Int64() != 0 {
					t.Fatal("Balance should be 0")
				}
			})

			t.Run("Buy if address is in participants list", func(t *testing.T) {
				addr := DeployContract(user[0], sim, big.NewInt(time.Now().Unix()+2), big.NewInt(stageDuration), big.NewInt(stageDuration), big.NewInt(stageDuration), user[0].Opts.From, user[0].Opts.From)
				//owner session
				session := GetCrowdsaleSession(user[0], sim, addr)
				if session == nil {
					t.Fatal("Should successfuly get contract session")
				}
				tokenSession := GetTokenSession(user[0], sim, addr)
				if tokenSession == nil {
					t.Fatal("Token session should not be nil")
				}

				tx, err := session.PrepareCrowdsale()
				if err != nil {
					t.Fatal("Failed to send TX prepareCrowdsale")
				}
				receipt := GetReceipt(tx, sim)
				if receipt.Status != 1 {
					t.Fatal("TX should not fail")
				}

				tx, err = session.AddPrivateParticipant(user[1].Opts.From)
				if err != nil {
					t.Fatal("TX should have been successfully created")
				}
				receipt = GetReceipt(tx, sim)
				if receipt.Status != 1 {
					t.Fatal("TX should be successful")
				}

				time.Sleep(3 * time.Second)

				tx = sendTransaction(user[1], sim, addr, "1000000000000000000")
				if tx == nil {
					t.Fatal("Transaction should have been successfully created")
				}
				receipt = GetReceipt(tx, sim)
				if receipt.Status != 1 {
					t.Fatal("TX should be successful")
				}

				balance, err := tokenSession.BalanceOf(user[1].Opts.From)
				if err != nil {
					t.Fatal("Transaction should have been successfully created")
				}
				if balance.Int64() != 0 {
					t.Fatal("Balance should be 0")
				}
			})
		})
	})

	t.Run("", func(t *testing.T) {

	})
}
