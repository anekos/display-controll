# To get models

```
# ddcutil detect
Display 1
   I2C bus:  /dev/i2c-3
   EDID synopsis:
      Mfg id:               DEL - Dell Inc.
      Model:                DELL U2718Q
      Product code:         64206  (0xface)
      Serial number:        DELL4KYASUII
      Binary serial number: 3735928559 (0xDEADBEEF)
      Manufacture year:     2017,  Week: 23
   VCP version:         2.1
....
```

# To get codes

```
# ddcutil --model 'DELL U2720QM' getvcp 0xdc
VCP code 0xdc (Display Mode                  ): Games (sl=0x05)
```
