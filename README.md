# Safe cash service BE

## Overview:

For purpose of handling consumers info and videos

## Architect:

Architect of all over system:

![Architect](./doc/Architect.jpg)

## Database design:

![Database](./doc/Database.jpg)

## Local dep:

For porpose of education, you can check more detail in folder setup with docker and docker-compose for local dep

## Production env to use:

`url`: `http://35.198.235.51:5000/`

## APIs:

**1. Register:**

`url`: `/apis/register`

`method`: `POST`

`Header`: 

```json
{
    "Content-Type": "application/json"
}
```

`Request`:

```json
{
    "email": "trungtin2qn1@gmail.com",
	"password": "1234567",
	"store_name": "Tin Huynh 1"
}
```

`Response`:

```json
{
    "id": "1544916627202511873",
    "email": "trungtin2qn1@gmail.com",
    "token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjE1NDQ5MTY2MjcyMDI1MTE4NzMiLCJlbWFpbCI6InRydW5ndGluMnFuMUBnbWFpbC5jb20iLCJzdG9yZV9pZCI6IjE1NDQ5MTY2MjcyMDI1MTE4NzMiLCJleHAiOjE1ODczOTI5NzAsImlzcyI6IlRpbiJ9.gK27j2jwyND7fi8uonD9u58X2jE9MiRAyz59FaQWoc1mcUS4kEniksKZxh4M3OfYnDinCBtjobfb0KTUEsj-7oQrixq2mf4p3hVTK8v1FCN50N6s0N2VDfgDFZ52rt5pnn6ugoBXjhUtdHUxVc8Scn7sLsSMmEyllYGw7Nog97eebBEx3lra5LlBQcGyv01nXi7pLvW4e4i4kq-0UKB9kVgfT4sbsFgiZiG7J5SLd11QX1T4PF0rpjidgy87Mm3R92aNfEZgu3eQcMW5yzvq3bHV6gEfxlGCY7zn8-3hESNxED4t6-Nsc61sDoEHQZhqkWCBE_KV5gmRvo1O_Ujz-siadxR30ToxJK7xl8MAhqjnpu_x9_bSrOAUEn3EX4PA7LYsWwYuA1ZVX-wejI4TyV7dJUz8-OBwyL7tLwQvU_IhuVGtrqU7cOacUOxSto4AD3R6RZJXmxjPWdWSpPQXChXUcb5jW-Dn2ZY9lYSn6f6OAV29btY8WKRTNi_wd9dBuvhzYSObYGTOzcqr2ExgXRO22S9_5mOqqc1i-tuhza0GOSI3_vViR8ZGWHmxNr-pJeVzXJ0PFKZ8IIKtQ7PVME1s_r1xDvUkKonhOqGPJalw5FxazmXHifQxrLb9LI1BkvZfl8fPwdPVD0Iazuk-6I1s6abHJANL126yTI91HaQ",
    "store_id": "1544916627202511873",
    "created_at": "2020-04-19T14:29:30.449653676Z",
    "updated_at": "2020-04-19T14:29:30.449653676Z"
}
```

**2. Login: **

`url`: `/apis/login`

`method`: `POST`

`Header`:

```json
{
    "Content-Type": "application/json"
}
```

`Request`:

```json
{
	"email": "trungtin2qn1@gmail.com",
	"password": "1234567",
	"store_name": "Tin Huynh 1"
}
```

`Response`:

```json
{
    "id": "1544916627202511873",
    "email": "trungtin2qn1@gmail.com",
    "token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjE1NDQ5MTY2MjcyMDI1MTE4NzMiLCJlbWFpbCI6InRydW5ndGluMnFuMUBnbWFpbC5jb20iLCJleHAiOjE1ODc0MjkzNjMsImlzcyI6IlRpbiJ9.Sq3sJkrvU44MFWP3cDTXgpYnQ7mF7T6DrXZMbtFpLEWJNa001-XeDsob48GO4zUfnukWl-nfz21q9p1S3zMn3-ts1-BpWQC40X_Jsg5li4YguvDtFx5FylUMI4a8k4QySQ76_8YuwH9zFcGM8WzBR6C7opwohPg0LGTvQccdyq-qknlcf0qsVxHJ1xm7eyQ1LwzX6TciBxqrpsbgNlyg96MZ6dJ-7bJWVjinvMqdN8XBc64P2rMr_mpdcxslxrft4xUFEG1O7JJyePQJ5j7WDrmx-7oE2nZ9kpu3uHGCCwxuBXZYoCuttoxBnYXwG7Opa7DRo1byr4WDn5NP4AdqTYxW_3LCwD7LMCprRaKoM9cMOIjCtRUoOAvHChr3OEcC_t5b135vsD27g4P6lxXNaCeVR8CODtIyLjHS8g5IJG3h2RznlerHK0fCftCX30y21rnttArBzV59aUiQkJDBv6M4LIpT8t-FBFPThdaPx4XvFUDq44_cVct4LdxTbjr5MQcP7s3xvhiyX46M3TV_6NO26Q3WJ3u0okFUY2EeLoZhtGFM9sxMVcFB0GH7csLrRx9HnxLnuFJ0xr_t74us77MQQOE0Whf3Ii1qQ9RRfRnTCPSvfJPL4dt_xgIxlf1bPqK2MTcpTd9fRdfyjrAVayweCnKTdRPTBVO5sCCD76c",
    "store_id": "1544916627202511873",
    "created_at": "2020-04-19T14:29:30.449654Z",
    "updated_at": "2020-04-19T14:29:30.449654Z"
}
```

**3. Upload media content:**

`url`: `/apis/auth/media`

`method`: `POST`

`Header`: 

```json
{
    "Authorization": "Tin eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjE1NDQ5MTY2MjcyMDI1MTE4NzMiLCJlbWFpbCI6InRydW5ndGluMnFuMUBnbWFpbC5jb20iLCJleHAiOjE1ODc0MjkzNjMsImlzcyI6IlRpbiJ9.Sq3sJkrvU44MFWP3cDTXgpYnQ7mF7T6DrXZMbtFpLEWJNa001-XeDsob48GO4zUfnukWl-nfz21q9p1S3zMn3-ts1-BpWQC40X_Jsg5li4YguvDtFx5FylUMI4a8k4QySQ76_8YuwH9zFcGM8WzBR6C7opwohPg0LGTvQccdyq-qknlcf0qsVxHJ1xm7eyQ1LwzX6TciBxqrpsbgNlyg96MZ6dJ-7bJWVjinvMqdN8XBc64P2rMr_mpdcxslxrft4xUFEG1O7JJyePQJ5j7WDrmx-7oE2nZ9kpu3uHGCCwxuBXZYoCuttoxBnYXwG7Opa7DRo1byr4WDn5NP4AdqTYxW_3LCwD7LMCprRaKoM9cMOIjCtRUoOAvHChr3OEcC_t5b135vsD27g4P6lxXNaCeVR8CODtIyLjHS8g5IJG3h2RznlerHK0fCftCX30y21rnttArBzV59aUiQkJDBv6M4LIpT8t-FBFPThdaPx4XvFUDq44_cVct4LdxTbjr5MQcP7s3xvhiyX46M3TV_6NO26Q3WJ3u0okFUY2EeLoZhtGFM9sxMVcFB0GH7csLrRx9HnxLnuFJ0xr_t74us77MQQOE0Whf3Ii1qQ9RRfRnTCPSvfJPL4dt_xgIxlf1bPqK2MTcpTd9fRdfyjrAVayweCnKTdRPTBVO5sCCD76c",
    "Content-Type": "multipart/form-data; boundary=<calculated when request is sent>"
}
```

`Request`:

```json
{
    "file": binary file
}
```

`Response`:

```json
{
    "message": "Success",
}
```

**4. Save notification token:**

`url`: `/apis/auth/notification/token`

`Method`: `POST`

`Header`:

```json
{
    "Content-Type": "application/json",
    "Authorization": "Tin eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjE1NDQ5MTY2MjcyMDI1MTE4NzMiLCJlbWFpbCI6InRydW5ndGluMnFuMUBnbWFpbC5jb20iLCJleHAiOjE1ODc0MjkzNjMsImlzcyI6IlRpbiJ9.Sq3sJkrvU44MFWP3cDTXgpYnQ7mF7T6DrXZMbtFpLEWJNa001-XeDsob48GO4zUfnukWl-nfz21q9p1S3zMn3-ts1-BpWQC40X_Jsg5li4YguvDtFx5FylUMI4a8k4QySQ76_8YuwH9zFcGM8WzBR6C7opwohPg0LGTvQccdyq-qknlcf0qsVxHJ1xm7eyQ1LwzX6TciBxqrpsbgNlyg96MZ6dJ-7bJWVjinvMqdN8XBc64P2rMr_mpdcxslxrft4xUFEG1O7JJyePQJ5j7WDrmx-7oE2nZ9kpu3uHGCCwxuBXZYoCuttoxBnYXwG7Opa7DRo1byr4WDn5NP4AdqTYxW_3LCwD7LMCprRaKoM9cMOIjCtRUoOAvHChr3OEcC_t5b135vsD27g4P6lxXNaCeVR8CODtIyLjHS8g5IJG3h2RznlerHK0fCftCX30y21rnttArBzV59aUiQkJDBv6M4LIpT8t-FBFPThdaPx4XvFUDq44_cVct4LdxTbjr5MQcP7s3xvhiyX46M3TV_6NO26Q3WJ3u0okFUY2EeLoZhtGFM9sxMVcFB0GH7csLrRx9HnxLnuFJ0xr_t74us77MQQOE0Whf3Ii1qQ9RRfRnTCPSvfJPL4dt_xgIxlf1bPqK2MTcpTd9fRdfyjrAVayweCnKTdRPTBVO5sCCD76c",
}
```

`Query`:

```json
{
    "offset": 0,
    "limit": 0,
    "type": "store" (for get all store notifications)
}
```

`Request`:

```json
{
	"token": "123"
}
```

`Response`:

```json
{
    "id": "1545233875968463873",
    "user_id": "1544916627202511873",
    "created_at": "2020-04-20T00:59:49.357387251Z",
    "updated_at": "2020-04-20T00:59:49.357387251Z"
}
```

**5. Send notification:**

`url`: `/apis/auth/notification`

`Method`: `POST`

`Header`: 

```json
{
    "Content-Type": "application/json",
    "Authorization": "Tin eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjE1NDQ5MTY2MjcyMDI1MTE4NzMiLCJlbWFpbCI6InRydW5ndGluMnFuMUBnbWFpbC5jb20iLCJleHAiOjE1ODc0MjkzNjMsImlzcyI6IlRpbiJ9.Sq3sJkrvU44MFWP3cDTXgpYnQ7mF7T6DrXZMbtFpLEWJNa001-XeDsob48GO4zUfnukWl-nfz21q9p1S3zMn3-ts1-BpWQC40X_Jsg5li4YguvDtFx5FylUMI4a8k4QySQ76_8YuwH9zFcGM8WzBR6C7opwohPg0LGTvQccdyq-qknlcf0qsVxHJ1xm7eyQ1LwzX6TciBxqrpsbgNlyg96MZ6dJ-7bJWVjinvMqdN8XBc64P2rMr_mpdcxslxrft4xUFEG1O7JJyePQJ5j7WDrmx-7oE2nZ9kpu3uHGCCwxuBXZYoCuttoxBnYXwG7Opa7DRo1byr4WDn5NP4AdqTYxW_3LCwD7LMCprRaKoM9cMOIjCtRUoOAvHChr3OEcC_t5b135vsD27g4P6lxXNaCeVR8CODtIyLjHS8g5IJG3h2RznlerHK0fCftCX30y21rnttArBzV59aUiQkJDBv6M4LIpT8t-FBFPThdaPx4XvFUDq44_cVct4LdxTbjr5MQcP7s3xvhiyX46M3TV_6NO26Q3WJ3u0okFUY2EeLoZhtGFM9sxMVcFB0GH7csLrRx9HnxLnuFJ0xr_t74us77MQQOE0Whf3Ii1qQ9RRfRnTCPSvfJPL4dt_xgIxlf1bPqK2MTcpTd9fRdfyjrAVayweCnKTdRPTBVO5sCCD76c",
}
```

`Response`:

```json
{
    "message": "success"
}
```

**6. Add user for a store:**

`url`: `/apis/auth/register`

`Method`: `POST`

`Header`:

```json
{
    "Content-Type": "application/json",
    "Authorization": "Tin eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjE1NDQ5MTY2MjcyMDI1MTE4NzMiLCJlbWFpbCI6InRydW5ndGluMnFuMUBnbWFpbC5jb20iLCJleHAiOjE1ODc0MjkzNjMsImlzcyI6IlRpbiJ9.Sq3sJkrvU44MFWP3cDTXgpYnQ7mF7T6DrXZMbtFpLEWJNa001-XeDsob48GO4zUfnukWl-nfz21q9p1S3zMn3-ts1-BpWQC40X_Jsg5li4YguvDtFx5FylUMI4a8k4QySQ76_8YuwH9zFcGM8WzBR6C7opwohPg0LGTvQccdyq-qknlcf0qsVxHJ1xm7eyQ1LwzX6TciBxqrpsbgNlyg96MZ6dJ-7bJWVjinvMqdN8XBc64P2rMr_mpdcxslxrft4xUFEG1O7JJyePQJ5j7WDrmx-7oE2nZ9kpu3uHGCCwxuBXZYoCuttoxBnYXwG7Opa7DRo1byr4WDn5NP4AdqTYxW_3LCwD7LMCprRaKoM9cMOIjCtRUoOAvHChr3OEcC_t5b135vsD27g4P6lxXNaCeVR8CODtIyLjHS8g5IJG3h2RznlerHK0fCftCX30y21rnttArBzV59aUiQkJDBv6M4LIpT8t-FBFPThdaPx4XvFUDq44_cVct4LdxTbjr5MQcP7s3xvhiyX46M3TV_6NO26Q3WJ3u0okFUY2EeLoZhtGFM9sxMVcFB0GH7csLrRx9HnxLnuFJ0xr_t74us77MQQOE0Whf3Ii1qQ9RRfRnTCPSvfJPL4dt_xgIxlf1bPqK2MTcpTd9fRdfyjrAVayweCnKTdRPTBVO5sCCD76c",
}
```

`Request`:

```json
{
    "email": string,
    "password": string
}
```

`Response`:

```json
{
    "id": "1544916627202511873",
    "email": "trungtin2qn1@gmail.com",
    "token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjE1NDQ5MTY2MjcyMDI1MTE4NzMiLCJlbWFpbCI6InRydW5ndGluMnFuMUBnbWFpbC5jb20iLCJleHAiOjE1ODc0MjkzNjMsImlzcyI6IlRpbiJ9.Sq3sJkrvU44MFWP3cDTXgpYnQ7mF7T6DrXZMbtFpLEWJNa001-XeDsob48GO4zUfnukWl-nfz21q9p1S3zMn3-ts1-BpWQC40X_Jsg5li4YguvDtFx5FylUMI4a8k4QySQ76_8YuwH9zFcGM8WzBR6C7opwohPg0LGTvQccdyq-qknlcf0qsVxHJ1xm7eyQ1LwzX6TciBxqrpsbgNlyg96MZ6dJ-7bJWVjinvMqdN8XBc64P2rMr_mpdcxslxrft4xUFEG1O7JJyePQJ5j7WDrmx-7oE2nZ9kpu3uHGCCwxuBXZYoCuttoxBnYXwG7Opa7DRo1byr4WDn5NP4AdqTYxW_3LCwD7LMCprRaKoM9cMOIjCtRUoOAvHChr3OEcC_t5b135vsD27g4P6lxXNaCeVR8CODtIyLjHS8g5IJG3h2RznlerHK0fCftCX30y21rnttArBzV59aUiQkJDBv6M4LIpT8t-FBFPThdaPx4XvFUDq44_cVct4LdxTbjr5MQcP7s3xvhiyX46M3TV_6NO26Q3WJ3u0okFUY2EeLoZhtGFM9sxMVcFB0GH7csLrRx9HnxLnuFJ0xr_t74us77MQQOE0Whf3Ii1qQ9RRfRnTCPSvfJPL4dt_xgIxlf1bPqK2MTcpTd9fRdfyjrAVayweCnKTdRPTBVO5sCCD76c",
    "store_id": "1544916627202511873",
    "created_at": "2020-04-19T14:29:30.449654Z",
    "updated_at": "2020-04-19T14:29:30.449654Z"
}
```

***Note: For using token prefix need to be "Tin " example: "Tin 12312321....."***

**7. Get user info (self):**

`URL`: `/apis/auth/self`

`Method`: `GET`

`Header`: 

```json
{
    "Content-Type": "application/json",
    "Authorization": "Tin eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjE1NzkwMzQzNjI2MTgzMTk4NzMiLCJlbWFpbCI6InRydW5ndGluMnFuMUBnbWFpbC5jb20iLCJzdG9yZV9pZCI6IjE1NzkwMzQzNjI2MTgzMTk4NzUiLCJleHAiOjE1OTE0NjAxMjEsImlzcyI6IlRpbiJ9.cT_Hycmc__2hD4nKGZrPYW9F3BwIa5PTok5EaHQWXo6mbwoqP4hJajckxujVpftSc36t7qGkvRvpFDe_ejk9tX_4COLe6HBx81gVYE5dnRDWPlvgkdoAN4lN3Y5l3y5ERGnL0hXrKP54TsITl5p0BV4ys2IClS5mANPRPOKoCSYYUy_VvqI52yHV5JxVsOVy7HpMRl4JpEkArL0vNTWFGzrl6XvpiLO913PLd5Cw1p1k78AKo_cWVqMm_1j0pw5IMR1fuQtHH64FLIY9xO-UZww_j4bVMEFsIHGh0UopmKc7AWUnirGRakCDwJOCkzQ0JPV92YhDt1CtBIB6uQ0louFeNPtq7f5l7ni1q6Vt8z8tBo4YVDEKKRchXhrqm2XQ1vZBDOmpM8K_XMFUCq-twRgpv5sBZDE1ANu20Kal6yJUyUmcGKLmxveg6vHvE9oSzCBr380hfgrtrTeqd_Xdi45iqCuMhC8UeWKokmwwr-DykYyAnjtCnQms23PlmzVj8vIS5sfZVPEfmHbUDCioQEbvM7lw3Gr1Z2l2LASGP3-P5xQZ6FSqgSSHq-m7tNA76-6HvMljXh5yZrGtGDQ8Jv35ZeMKSAdo-kYdjTxOqiUZHtAl43bkInehCVFpL_glZg8l2cmo0n_87sTeoe41oQcizk_XyIeZCwtMD5iOj-8"
}
```

`Response`:

```json
{
    "id": "1579034362618319873",
    "email": "trungtin2qn1@gmail.com",
    "store_id": "1579034362618319875",
    "created_at": "2020-06-05T23:15:21.94383Z",
    "updated_at": "2020-06-06T00:29:10.506767Z"
}
```

**8. Update user info (self):**

`URL`: `/apis/auth/self`

`Method`: `PUT`

`Header`: 

```json
{
    "Content-Type": "application/json",
    "Authorization": "Tin eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjE1NzkwMzQzNjI2MTgzMTk4NzMiLCJlbWFpbCI6InRydW5ndGluMnFuMUBnbWFpbC5jb20iLCJzdG9yZV9pZCI6IjE1NzkwMzQzNjI2MTgzMTk4NzUiLCJleHAiOjE1OTE0NjAxMjEsImlzcyI6IlRpbiJ9.cT_Hycmc__2hD4nKGZrPYW9F3BwIa5PTok5EaHQWXo6mbwoqP4hJajckxujVpftSc36t7qGkvRvpFDe_ejk9tX_4COLe6HBx81gVYE5dnRDWPlvgkdoAN4lN3Y5l3y5ERGnL0hXrKP54TsITl5p0BV4ys2IClS5mANPRPOKoCSYYUy_VvqI52yHV5JxVsOVy7HpMRl4JpEkArL0vNTWFGzrl6XvpiLO913PLd5Cw1p1k78AKo_cWVqMm_1j0pw5IMR1fuQtHH64FLIY9xO-UZww_j4bVMEFsIHGh0UopmKc7AWUnirGRakCDwJOCkzQ0JPV92YhDt1CtBIB6uQ0louFeNPtq7f5l7ni1q6Vt8z8tBo4YVDEKKRchXhrqm2XQ1vZBDOmpM8K_XMFUCq-twRgpv5sBZDE1ANu20Kal6yJUyUmcGKLmxveg6vHvE9oSzCBr380hfgrtrTeqd_Xdi45iqCuMhC8UeWKokmwwr-DykYyAnjtCnQms23PlmzVj8vIS5sfZVPEfmHbUDCioQEbvM7lw3Gr1Z2l2LASGP3-P5xQZ6FSqgSSHq-m7tNA76-6HvMljXh5yZrGtGDQ8Jv35ZeMKSAdo-kYdjTxOqiUZHtAl43bkInehCVFpL_glZg8l2cmo0n_87sTeoe41oQcizk_XyIeZCwtMD5iOj-8"
}
```

`Response`:

```json:
{
    "id": "1579034362618319873",
    "email": "trungtin2qn1@gmail.com",
    "phone_number": "123456789",
    "first_name": "Tin",
    "last_name": "Huynh"
}
```

**9. Change password:**

`URL`: `/apis/auth/self/password`

`Method`: `PUT`

`Header`: 

```json
{
    "Content-Type": "application/json",
    "Authorization": "Tin eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjE1NzkwMzQzNjI2MTgzMTk4NzMiLCJlbWFpbCI6InRydW5ndGluMnFuMUBnbWFpbC5jb20iLCJzdG9yZV9pZCI6IjE1NzkwMzQzNjI2MTgzMTk4NzUiLCJleHAiOjE1OTE0NjAxMjEsImlzcyI6IlRpbiJ9.cT_Hycmc__2hD4nKGZrPYW9F3BwIa5PTok5EaHQWXo6mbwoqP4hJajckxujVpftSc36t7qGkvRvpFDe_ejk9tX_4COLe6HBx81gVYE5dnRDWPlvgkdoAN4lN3Y5l3y5ERGnL0hXrKP54TsITl5p0BV4ys2IClS5mANPRPOKoCSYYUy_VvqI52yHV5JxVsOVy7HpMRl4JpEkArL0vNTWFGzrl6XvpiLO913PLd5Cw1p1k78AKo_cWVqMm_1j0pw5IMR1fuQtHH64FLIY9xO-UZww_j4bVMEFsIHGh0UopmKc7AWUnirGRakCDwJOCkzQ0JPV92YhDt1CtBIB6uQ0louFeNPtq7f5l7ni1q6Vt8z8tBo4YVDEKKRchXhrqm2XQ1vZBDOmpM8K_XMFUCq-twRgpv5sBZDE1ANu20Kal6yJUyUmcGKLmxveg6vHvE9oSzCBr380hfgrtrTeqd_Xdi45iqCuMhC8UeWKokmwwr-DykYyAnjtCnQms23PlmzVj8vIS5sfZVPEfmHbUDCioQEbvM7lw3Gr1Z2l2LASGP3-P5xQZ6FSqgSSHq-m7tNA76-6HvMljXh5yZrGtGDQ8Jv35ZeMKSAdo-kYdjTxOqiUZHtAl43bkInehCVFpL_glZg8l2cmo0n_87sTeoe41oQcizk_XyIeZCwtMD5iOj-8"
}
```

`Request`:

```json
{
	"old_password": "1234567",
	"new_password": "12345678"
}
```

`Response:`

```json
{
    "message": "Success"
}
```

**10. Unlock smart withdrawal**

`URL`: `/apis/auth/unlock`

`Method`: `POST`

`Header`: 

```json
{
    "Content-Type": "application/json",
    "Authorization": "Tin eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjE1NzkwMzQzNjI2MTgzMTk4NzMiLCJlbWFpbCI6InRydW5ndGluMnFuMUBnbWFpbC5jb20iLCJzdG9yZV9pZCI6IjE1NzkwMzQzNjI2MTgzMTk4NzUiLCJleHAiOjE1OTE0NjAxMjEsImlzcyI6IlRpbiJ9.cT_Hycmc__2hD4nKGZrPYW9F3BwIa5PTok5EaHQWXo6mbwoqP4hJajckxujVpftSc36t7qGkvRvpFDe_ejk9tX_4COLe6HBx81gVYE5dnRDWPlvgkdoAN4lN3Y5l3y5ERGnL0hXrKP54TsITl5p0BV4ys2IClS5mANPRPOKoCSYYUy_VvqI52yHV5JxVsOVy7HpMRl4JpEkArL0vNTWFGzrl6XvpiLO913PLd5Cw1p1k78AKo_cWVqMm_1j0pw5IMR1fuQtHH64FLIY9xO-UZww_j4bVMEFsIHGh0UopmKc7AWUnirGRakCDwJOCkzQ0JPV92YhDt1CtBIB6uQ0louFeNPtq7f5l7ni1q6Vt8z8tBo4YVDEKKRchXhrqm2XQ1vZBDOmpM8K_XMFUCq-twRgpv5sBZDE1ANu20Kal6yJUyUmcGKLmxveg6vHvE9oSzCBr380hfgrtrTeqd_Xdi45iqCuMhC8UeWKokmwwr-DykYyAnjtCnQms23PlmzVj8vIS5sfZVPEfmHbUDCioQEbvM7lw3Gr1Z2l2LASGP3-P5xQZ6FSqgSSHq-m7tNA76-6HvMljXh5yZrGtGDQ8Jv35ZeMKSAdo-kYdjTxOqiUZHtAl43bkInehCVFpL_glZg8l2cmo0n_87sTeoe41oQcizk_XyIeZCwtMD5iOj-8"
}
```

`Request`:

```json
{
	"content": "Success in unlocking withdrawal",
	"is_success": true
}
```

`Response:`

```json
{
    "id": "1583213810393551873",
    "content": "Success in unlocking withdrawal",
    "is_success": true,
    "user_id": "1579034362618319873",
    "created_at": "2020-06-11T17:39:10.868472977+07:00",
    "updated_at": "2020-06-11T17:39:10.868472977+07:00"
}
```

**11. List all unlocking logs filter by user or store:**

`URL`: `/apis/auth/unlock/logs`

`Method`: `GET`

`Query`:

```json
{
    "store_id": "1583213810393551873"
}
```

`Header`: 

```json
{
    "Content-Type": "application/json",
    "Authorization": "Tin eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjE1NzkwMzQzNjI2MTgzMTk4NzMiLCJlbWFpbCI6InRydW5ndGluMnFuMUBnbWFpbC5jb20iLCJzdG9yZV9pZCI6IjE1NzkwMzQzNjI2MTgzMTk4NzUiLCJleHAiOjE1OTE0NjAxMjEsImlzcyI6IlRpbiJ9.cT_Hycmc__2hD4nKGZrPYW9F3BwIa5PTok5EaHQWXo6mbwoqP4hJajckxujVpftSc36t7qGkvRvpFDe_ejk9tX_4COLe6HBx81gVYE5dnRDWPlvgkdoAN4lN3Y5l3y5ERGnL0hXrKP54TsITl5p0BV4ys2IClS5mANPRPOKoCSYYUy_VvqI52yHV5JxVsOVy7HpMRl4JpEkArL0vNTWFGzrl6XvpiLO913PLd5Cw1p1k78AKo_cWVqMm_1j0pw5IMR1fuQtHH64FLIY9xO-UZww_j4bVMEFsIHGh0UopmKc7AWUnirGRakCDwJOCkzQ0JPV92YhDt1CtBIB6uQ0louFeNPtq7f5l7ni1q6Vt8z8tBo4YVDEKKRchXhrqm2XQ1vZBDOmpM8K_XMFUCq-twRgpv5sBZDE1ANu20Kal6yJUyUmcGKLmxveg6vHvE9oSzCBr380hfgrtrTeqd_Xdi45iqCuMhC8UeWKokmwwr-DykYyAnjtCnQms23PlmzVj8vIS5sfZVPEfmHbUDCioQEbvM7lw3Gr1Z2l2LASGP3-P5xQZ6FSqgSSHq-m7tNA76-6HvMljXh5yZrGtGDQ8Jv35ZeMKSAdo-kYdjTxOqiUZHtAl43bkInehCVFpL_glZg8l2cmo0n_87sTeoe41oQcizk_XyIeZCwtMD5iOj-8"
}
```

`Response:`

```json
[
    {
        "id": "1583213810393551873",
        "content": "Success in unlocking withdrawal",
        "is_success": true,
        "user_id": "1579034362618319873",
        "created_at": "2020-06-11T17:39:10.868473Z",
        "updated_at": "2020-06-11T17:39:10.868473Z"
    }
]
```

**12. Get all notifications (self):**

`URL`: `/apis/auth/unlock`

`Method`: `GET`

`Header`: 

```json
{
    "Content-Type": "application/json",
    "Authorization": "Tin eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjE1NzkwMzQzNjI2MTgzMTk4NzMiLCJlbWFpbCI6InRydW5ndGluMnFuMUBnbWFpbC5jb20iLCJzdG9yZV9pZCI6IjE1NzkwMzQzNjI2MTgzMTk4NzUiLCJleHAiOjE1OTE0NjAxMjEsImlzcyI6IlRpbiJ9.cT_Hycmc__2hD4nKGZrPYW9F3BwIa5PTok5EaHQWXo6mbwoqP4hJajckxujVpftSc36t7qGkvRvpFDe_ejk9tX_4COLe6HBx81gVYE5dnRDWPlvgkdoAN4lN3Y5l3y5ERGnL0hXrKP54TsITl5p0BV4ys2IClS5mANPRPOKoCSYYUy_VvqI52yHV5JxVsOVy7HpMRl4JpEkArL0vNTWFGzrl6XvpiLO913PLd5Cw1p1k78AKo_cWVqMm_1j0pw5IMR1fuQtHH64FLIY9xO-UZww_j4bVMEFsIHGh0UopmKc7AWUnirGRakCDwJOCkzQ0JPV92YhDt1CtBIB6uQ0louFeNPtq7f5l7ni1q6Vt8z8tBo4YVDEKKRchXhrqm2XQ1vZBDOmpM8K_XMFUCq-twRgpv5sBZDE1ANu20Kal6yJUyUmcGKLmxveg6vHvE9oSzCBr380hfgrtrTeqd_Xdi45iqCuMhC8UeWKokmwwr-DykYyAnjtCnQms23PlmzVj8vIS5sfZVPEfmHbUDCioQEbvM7lw3Gr1Z2l2LASGP3-P5xQZ6FSqgSSHq-m7tNA76-6HvMljXh5yZrGtGDQ8Jv35ZeMKSAdo-kYdjTxOqiUZHtAl43bkInehCVFpL_glZg8l2cmo0n_87sTeoe41oQcizk_XyIeZCwtMD5iOj-8"
}
```

`Response:`

```json
[
    {
        "id": "1583222308053455873",
        "is_read": false,
        "user_id": "1579034362618319873",
        "created_at": "2020-06-11T17:56:03.485181Z",
        "updated_at": "2020-06-11T17:56:03.485181Z"
    },
    {
        "id": "1583222668763599874",
        "is_read": false,
        "user_id": "1579034362618319873",
        "created_at": "2020-06-11T17:56:46.573809Z",
        "updated_at": "2020-06-11T17:56:46.573809Z"
    }
]
```

**13. Update notification status:**

`URL`: `/apis/auth/notification/:notification_id`

`Method`: `POST`

`Header`: 

```json
{
    "Content-Type": "application/json",
    "Authorization": "Tin eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjE1NzkwMzQzNjI2MTgzMTk4NzMiLCJlbWFpbCI6InRydW5ndGluMnFuMUBnbWFpbC5jb20iLCJzdG9yZV9pZCI6IjE1NzkwMzQzNjI2MTgzMTk4NzUiLCJleHAiOjE1OTE0NjAxMjEsImlzcyI6IlRpbiJ9.cT_Hycmc__2hD4nKGZrPYW9F3BwIa5PTok5EaHQWXo6mbwoqP4hJajckxujVpftSc36t7qGkvRvpFDe_ejk9tX_4COLe6HBx81gVYE5dnRDWPlvgkdoAN4lN3Y5l3y5ERGnL0hXrKP54TsITl5p0BV4ys2IClS5mANPRPOKoCSYYUy_VvqI52yHV5JxVsOVy7HpMRl4JpEkArL0vNTWFGzrl6XvpiLO913PLd5Cw1p1k78AKo_cWVqMm_1j0pw5IMR1fuQtHH64FLIY9xO-UZww_j4bVMEFsIHGh0UopmKc7AWUnirGRakCDwJOCkzQ0JPV92YhDt1CtBIB6uQ0louFeNPtq7f5l7ni1q6Vt8z8tBo4YVDEKKRchXhrqm2XQ1vZBDOmpM8K_XMFUCq-twRgpv5sBZDE1ANu20Kal6yJUyUmcGKLmxveg6vHvE9oSzCBr380hfgrtrTeqd_Xdi45iqCuMhC8UeWKokmwwr-DykYyAnjtCnQms23PlmzVj8vIS5sfZVPEfmHbUDCioQEbvM7lw3Gr1Z2l2LASGP3-P5xQZ6FSqgSSHq-m7tNA76-6HvMljXh5yZrGtGDQ8Jv35ZeMKSAdo-kYdjTxOqiUZHtAl43bkInehCVFpL_glZg8l2cmo0n_87sTeoe41oQcizk_XyIeZCwtMD5iOj-8"
}
```

`Param`:

```json
{
    "notification_id": "1583222308053455873"
}
```

`Response:`

```json
{
    "message": "Success"
}
```

**14.  Get all staffs in store (for merchant only)**

`URL`: `/apis/auth/staffs`

`Method`: `GET`

`Header`: 

```json
{
    "Content-Type": "application/json",
    "Authorization": "Tin eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjE1NzkwMzQzNjI2MTgzMTk4NzMiLCJlbWFpbCI6InRydW5ndGluMnFuMUBnbWFpbC5jb20iLCJzdG9yZV9pZCI6IjE1NzkwMzQzNjI2MTgzMTk4NzUiLCJleHAiOjE1OTE0NjAxMjEsImlzcyI6IlRpbiJ9.cT_Hycmc__2hD4nKGZrPYW9F3BwIa5PTok5EaHQWXo6mbwoqP4hJajckxujVpftSc36t7qGkvRvpFDe_ejk9tX_4COLe6HBx81gVYE5dnRDWPlvgkdoAN4lN3Y5l3y5ERGnL0hXrKP54TsITl5p0BV4ys2IClS5mANPRPOKoCSYYUy_VvqI52yHV5JxVsOVy7HpMRl4JpEkArL0vNTWFGzrl6XvpiLO913PLd5Cw1p1k78AKo_cWVqMm_1j0pw5IMR1fuQtHH64FLIY9xO-UZww_j4bVMEFsIHGh0UopmKc7AWUnirGRakCDwJOCkzQ0JPV92YhDt1CtBIB6uQ0louFeNPtq7f5l7ni1q6Vt8z8tBo4YVDEKKRchXhrqm2XQ1vZBDOmpM8K_XMFUCq-twRgpv5sBZDE1ANu20Kal6yJUyUmcGKLmxveg6vHvE9oSzCBr380hfgrtrTeqd_Xdi45iqCuMhC8UeWKokmwwr-DykYyAnjtCnQms23PlmzVj8vIS5sfZVPEfmHbUDCioQEbvM7lw3Gr1Z2l2LASGP3-P5xQZ6FSqgSSHq-m7tNA76-6HvMljXh5yZrGtGDQ8Jv35ZeMKSAdo-kYdjTxOqiUZHtAl43bkInehCVFpL_glZg8l2cmo0n_87sTeoe41oQcizk_XyIeZCwtMD5iOj-8"
}
```

`Response:`

```json
[
    {
        "id": "1579034362618319873",
        "email": "trungtin2qn1@gmail.com",
        "phone_number": "123456789",
        "first_name": "Tin",
        "last_name": "Huynh",
        "store_id": "1579034362618319875",
        "created_at": "2020-06-05T23:15:21.94383Z",
        "updated_at": "2020-06-11T17:31:21.107445Z"
    }
]
```

**15. Upload media file from services:**

`URL`: `/services/stores/{store_id}/medias`

`Method`: `GET`

`Header`: 

```json
{
    "store_credential": "test",
    "Content-Type": "multipart/form-data"
}
```

`Request`:

```json
{
    "screenshot": binary file (image),
    "video": binary file (video),
    "user_id": "1620448417",
    "unlocking_log_id": "1620448417"
}
```

`Response:`

```json
{
    "message": string
}
```


**16. Get medias file by unlocking id**

`URL`: `/apis/auth/store/{store_id}/unlock/{unlock_id}/medias`

`Method`: `GET`

`Param`: 

```json
{
    "store_id": anything,
    "unlock_id": "12312"
}
```

`Header`: 

```json
{
    "Authorization": "Tin eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjE1NzkwMzQzNjI2MTgzMTk4NzMiLCJlbWFpbCI6InRydW5ndGluMnFuMUBnbWFpbC5jb20iLCJzdG9yZV9pZCI6IjE1NzkwMzQzNjI2MTgzMTk4NzUiLCJleHAiOjE1OTE0NjAxMjEsImlzcyI6IlRpbiJ9.cT_Hycmc__2hD4nKGZrPYW9F3BwIa5PTok5EaHQWXo6mbwoqP4hJajckxujVpftSc36t7qGkvRvpFDe_ejk9tX_4COLe6HBx81gVYE5dnRDWPlvgkdoAN4lN3Y5l3y5ERGnL0hXrKP54TsITl5p0BV4ys2IClS5mANPRPOKoCSYYUy_VvqI52yHV5JxVsOVy7HpMRl4JpEkArL0vNTWFGzrl6XvpiLO913PLd5Cw1p1k78AKo_cWVqMm_1j0pw5IMR1fuQtHH64FLIY9xO-UZww_j4bVMEFsIHGh0UopmKc7AWUnirGRakCDwJOCkzQ0JPV92YhDt1CtBIB6uQ0louFeNPtq7f5l7ni1q6Vt8z8tBo4YVDEKKRchXhrqm2XQ1vZBDOmpM8K_XMFUCq-twRgpv5sBZDE1ANu20Kal6yJUyUmcGKLmxveg6vHvE9oSzCBr380hfgrtrTeqd_Xdi45iqCuMhC8UeWKokmwwr-DykYyAnjtCnQms23PlmzVj8vIS5sfZVPEfmHbUDCioQEbvM7lw3Gr1Z2l2LASGP3-P5xQZ6FSqgSSHq-m7tNA76-6HvMljXh5yZrGtGDQ8Jv35ZeMKSAdo-kYdjTxOqiUZHtAl43bkInehCVFpL_glZg8l2cmo0n_87sTeoe41oQcizk_XyIeZCwtMD5iOj-8"
}
```

`Response:`

```json
[
    {
        "id": "1621330090",
        "name": "http://35.240.249.239:5000/public/e8171cb5-3b3e-435c-8f20-0fa0c2e7e6a3.jpg",
        "type": "image",
        "thumbnail": "e8171cb5-3b3e-435c-8f20-0fa0c2e7e6a3.jpg",
        "user_id": "1620869281",
        "store_id": "1620869283",
        "created_at": "2020-08-03T16:31:48.281135Z",
        "updated_at": "2020-08-03T16:31:48.281135Z"
    },
    {
        "id": "1621330091",
        "name": "http://35.240.249.239:5000/public/2cc23652-c633-4e9b-8f77-47596f0d5ace.png",
        "type": "thumbnail_video",
        "thumbnail": "2cc23652-c633-4e9b-8f77-47596f0d5ace.png",
        "user_id": "1620869281",
        "store_id": "1620869283",
        "created_at": "2020-08-03T16:31:50.564507Z",
        "updated_at": "2020-08-03T16:31:50.564507Z"
    },
    {
        "id": "1621330092",
        "name": "http://35.240.249.239:5000/public/41abaa06-c665-41e5-91a2-440f0d53d2e7.mp4",
        "type": "video",
        "thumbnail": "2cc23652-c633-4e9b-8f77-47596f0d5ace.png",
        "user_id": "1620869281",
        "store_id": "1620869283",
        "created_at": "2020-08-03T16:31:50.646623Z",
        "updated_at": "2020-08-03T16:31:50.646623Z"
    }
]
```

**17. Get medias by store id**

`URL`: `/apis/auth/store/{store_id}/medias`

`Method`: `GET`

`Param`: 

```json
{
    "store_id": anything
}
```

`Header`: 

```json
{
    "Authorization": "Tin eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjE1NzkwMzQzNjI2MTgzMTk4NzMiLCJlbWFpbCI6InRydW5ndGluMnFuMUBnbWFpbC5jb20iLCJzdG9yZV9pZCI6IjE1NzkwMzQzNjI2MTgzMTk4NzUiLCJleHAiOjE1OTE0NjAxMjEsImlzcyI6IlRpbiJ9.cT_Hycmc__2hD4nKGZrPYW9F3BwIa5PTok5EaHQWXo6mbwoqP4hJajckxujVpftSc36t7qGkvRvpFDe_ejk9tX_4COLe6HBx81gVYE5dnRDWPlvgkdoAN4lN3Y5l3y5ERGnL0hXrKP54TsITl5p0BV4ys2IClS5mANPRPOKoCSYYUy_VvqI52yHV5JxVsOVy7HpMRl4JpEkArL0vNTWFGzrl6XvpiLO913PLd5Cw1p1k78AKo_cWVqMm_1j0pw5IMR1fuQtHH64FLIY9xO-UZww_j4bVMEFsIHGh0UopmKc7AWUnirGRakCDwJOCkzQ0JPV92YhDt1CtBIB6uQ0louFeNPtq7f5l7ni1q6Vt8z8tBo4YVDEKKRchXhrqm2XQ1vZBDOmpM8K_XMFUCq-twRgpv5sBZDE1ANu20Kal6yJUyUmcGKLmxveg6vHvE9oSzCBr380hfgrtrTeqd_Xdi45iqCuMhC8UeWKokmwwr-DykYyAnjtCnQms23PlmzVj8vIS5sfZVPEfmHbUDCioQEbvM7lw3Gr1Z2l2LASGP3-P5xQZ6FSqgSSHq-m7tNA76-6HvMljXh5yZrGtGDQ8Jv35ZeMKSAdo-kYdjTxOqiUZHtAl43bkInehCVFpL_glZg8l2cmo0n_87sTeoe41oQcizk_XyIeZCwtMD5iOj-8"
}
```

`Response:`

```json
[
    {
        "images": [
            {
                "url": "ef4052c9-8072-4fb9-8da7-4d17e95527b4.jpg",
                "type": "image"
            },
            {
                "url": "a170c219-d4b8-4cb5-a22b-09416a31e649.png",
                "type": "thumbnail_video"
            },
            {
                "url": "27b9bf92-0084-4fc3-b577-e64377ca967f.jpg",
                "type": "image"
            },
            {
                "url": "e6efa47c-3dac-4ea5-a0f4-36b8d2d8600d.png",
                "type": "thumbnail_video"
            },
            {
                "url": "2cbcb108-0927-4a40-98b1-2fab11611d0a.jpg",
                "type": "image"
            },
            {
                "url": "28621fae-bdb9-4d9e-93cb-be0095eebac3.png",
                "type": "thumbnail_video"
            },
            {
                "url": "ee044430-dfe9-4b26-8fa6-3763970f69d2.jpg",
                "type": "image"
            },
            {
                "url": "11b4e069-5977-4f84-9356-4f33ab82a065.png",
                "type": "thumbnail_video"
            },
            {
                "url": "99ee8a5d-0f3c-4f6b-a819-3954d22d6030.jpg",
                "type": "image"
            },
            {
                "url": "db609b12-0461-4e69-a45c-697dda5c6b10.png",
                "type": "thumbnail_video"
            },
            {
                "url": "79c186a0-df4b-4e4b-b626-a7a2cf4e13c1.jpg",
                "type": "image"
            },
            {
                "url": "cd7b81cb-90f0-4d52-b83f-b642508e3ecf.png",
                "type": "thumbnail_video"
            },
            {
                "url": "fe90b1b9-4d7a-4213-b0af-23c4309bb6af.jpg",
                "type": "image"
            },
            {
                "url": "b89d5263-75ac-4847-bd93-bdeeac27d65f.png",
                "type": "thumbnail_video"
            },
            {
                "url": "ad7dab37-82e6-4de5-b92f-c4f950b438df.jpg",
                "type": "image"
            },
            {
                "url": "e43e0795-d56b-46e0-b568-f8d6fada2af2.png",
                "type": "thumbnail_video"
            },
            {
                "url": "cf3d35e6-d290-4629-9f90-b7b0033ba3f4.jpg",
                "type": "image"
            },
            {
                "url": "af01881f-9d76-4d62-b062-213762dac236.jpg",
                "type": "image"
            },
            {
                "url": "dce46d9c-0156-42cb-8e6e-846ce29a21bb.jpg",
                "type": "image"
            }
        ],
        "videos": [
            {
                "url": "e6343a50-d309-4e4c-9432-9e050fe44b81.mp4"
            },
            {
                "url": "d110538d-3b94-4a6c-a39b-94ebc1fd1512.mp4"
            },
            {
                "url": "734b205a-06d7-4485-9cfb-281e49712890.mp4"
            },
            {
                "url": "d3cb697e-f75f-4dba-8069-d8c240e82ed4.mp4"
            },
            {
                "url": "9a0e5637-075b-4461-9a0f-b9b4a0c71144.mp4"
            },
            {
                "url": "e2e9c186-4e3b-4a83-8cd0-d532cf567260.mp4"
            },
            {
                "url": "caf9cb2a-cbb3-427e-a89a-47f44c214f25.mp4"
            },
            {
                "url": "c0c71976-0940-4b38-8109-223c817307bc.mp4"
            },
            {
                "url": "62fb7a66-ae27-4c7a-b19a-c0456f784623.mp4"
            },
            {
                "url": "7688b833-0833-4753-8060-1e15011dea5c.mp4"
            },
            {
                "url": "792e76a3-5d8a-4b8d-96ef-299c37341c01.mp4"
            }
        ],
        "is_success": true,
        "username": "Tin Huynh",
        "created_at": "2020-07-30T18:22:54.42639Z",
        "updated_at": "2020-07-30T18:22:54.42639Z",
        "method": "face"
    },
    {
        "images": [],
        "videos": [],
        "is_success": true,
        "username": "Tin Huynh",
        "created_at": "2020-07-29T18:22:50.334872Z",
        "updated_at": "2020-07-29T18:22:50.334872Z",
        "method": "face"
    },
    {
        "images": [],
        "videos": [],
        "is_success": true,
        "username": "Tin Huynh",
        "created_at": "2020-07-29T18:13:01.321475Z",
        "updated_at": "2020-07-29T18:13:01.321475Z",
        "method": "face"
    },
    {
        "images": [],
        "videos": [],
        "is_success": true,
        "username": "Tin Huynh",
        "created_at": "2020-07-27T15:46:53.981802Z",
        "updated_at": "2020-07-27T15:46:53.981802Z",
        "method": "face"
    },
    {
        "images": [],
        "videos": [],
        "is_success": true,
        "username": "Tin Huynh",
        "created_at": "2020-07-26T23:25:07.096755Z",
        "updated_at": "2020-07-26T23:25:07.096755Z",
        "method": ""
    },
    {
        "images": [],
        "videos": [],
        "is_success": true,
        "username": "Tin Huynh",
        "created_at": "2020-07-26T12:47:11.521552Z",
        "updated_at": "2020-07-26T12:47:11.521552Z",
        "method": "face"
    },
    {
        "images": [],
        "videos": [],
        "is_success": true,
        "username": "Tin Huynh",
        "created_at": "2020-07-26T12:32:14.99194Z",
        "updated_at": "2020-07-26T12:32:14.99194Z",
        "method": "face"
    },
    {
        "images": [],
        "videos": [],
        "is_success": true,
        "username": "Tin Huynh",
        "created_at": "2020-07-26T12:31:58.711267Z",
        "updated_at": "2020-07-26T12:31:58.711267Z",
        "method": "face"
    }
]
```

**18.Update user info (with avatar)**

`URL`: `/apis/v1/auth/self`

`Method`: `PUT`

`Header`: 

```json
{
    "Content-Type": "multipart/form-data; boundary=<calculated when request is sent>",
    "Authorization": "Tin eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjE1NzkwMzQzNjI2MTgzMTk4NzMiLCJlbWFpbCI6InRydW5ndGluMnFuMUBnbWFpbC5jb20iLCJzdG9yZV9pZCI6IjE1NzkwMzQzNjI2MTgzMTk4NzUiLCJleHAiOjE1OTE0NjAxMjEsImlzcyI6IlRpbiJ9.cT_Hycmc__2hD4nKGZrPYW9F3BwIa5PTok5EaHQWXo6mbwoqP4hJajckxujVpftSc36t7qGkvRvpFDe_ejk9tX_4COLe6HBx81gVYE5dnRDWPlvgkdoAN4lN3Y5l3y5ERGnL0hXrKP54TsITl5p0BV4ys2IClS5mANPRPOKoCSYYUy_VvqI52yHV5JxVsOVy7HpMRl4JpEkArL0vNTWFGzrl6XvpiLO913PLd5Cw1p1k78AKo_cWVqMm_1j0pw5IMR1fuQtHH64FLIY9xO-UZww_j4bVMEFsIHGh0UopmKc7AWUnirGRakCDwJOCkzQ0JPV92YhDt1CtBIB6uQ0louFeNPtq7f5l7ni1q6Vt8z8tBo4YVDEKKRchXhrqm2XQ1vZBDOmpM8K_XMFUCq-twRgpv5sBZDE1ANu20Kal6yJUyUmcGKLmxveg6vHvE9oSzCBr380hfgrtrTeqd_Xdi45iqCuMhC8UeWKokmwwr-DykYyAnjtCnQms23PlmzVj8vIS5sfZVPEfmHbUDCioQEbvM7lw3Gr1Z2l2LASGP3-P5xQZ6FSqgSSHq-m7tNA76-6HvMljXh5yZrGtGDQ8Jv35ZeMKSAdo-kYdjTxOqiUZHtAl43bkInehCVFpL_glZg8l2cmo0n_87sTeoe41oQcizk_XyIeZCwtMD5iOj-8"
}
```

`Request:`

```yml
    "phone_number": string,
    "fist_name": string,
    "last_name": string,
    "avatar": image binary file
```

`Response:`

```json
{
    "id": "1621330090",
    "name": "http://35.240.249.239:5000/public/e8171cb5-3b3e-435c-8f20-0fa0c2e7e6a3.jpg",
    "type": "image",
    "thumbnail": "e8171cb5-3b3e-435c-8f20-0fa0c2e7e6a3.jpg",
    "user_id": "1620869281",
    "store_id": "1620869283",
    "created_at": "2020-08-03T16:31:48.281135Z",
    "updated_at": "2020-08-03T16:31:48.281135Z"
}
```

**19.Update avatar**

`URL`: `/apis/v1/auth/self/avatar`

`Method`: `PUT`

`Header`: 

```json
{
    "Content-Type": "multipart/form-data; boundary=<calculated when request is sent>",
    "Authorization": "Tin eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjE1NzkwMzQzNjI2MTgzMTk4NzMiLCJlbWFpbCI6InRydW5ndGluMnFuMUBnbWFpbC5jb20iLCJzdG9yZV9pZCI6IjE1NzkwMzQzNjI2MTgzMTk4NzUiLCJleHAiOjE1OTE0NjAxMjEsImlzcyI6IlRpbiJ9.cT_Hycmc__2hD4nKGZrPYW9F3BwIa5PTok5EaHQWXo6mbwoqP4hJajckxujVpftSc36t7qGkvRvpFDe_ejk9tX_4COLe6HBx81gVYE5dnRDWPlvgkdoAN4lN3Y5l3y5ERGnL0hXrKP54TsITl5p0BV4ys2IClS5mANPRPOKoCSYYUy_VvqI52yHV5JxVsOVy7HpMRl4JpEkArL0vNTWFGzrl6XvpiLO913PLd5Cw1p1k78AKo_cWVqMm_1j0pw5IMR1fuQtHH64FLIY9xO-UZww_j4bVMEFsIHGh0UopmKc7AWUnirGRakCDwJOCkzQ0JPV92YhDt1CtBIB6uQ0louFeNPtq7f5l7ni1q6Vt8z8tBo4YVDEKKRchXhrqm2XQ1vZBDOmpM8K_XMFUCq-twRgpv5sBZDE1ANu20Kal6yJUyUmcGKLmxveg6vHvE9oSzCBr380hfgrtrTeqd_Xdi45iqCuMhC8UeWKokmwwr-DykYyAnjtCnQms23PlmzVj8vIS5sfZVPEfmHbUDCioQEbvM7lw3Gr1Z2l2LASGP3-P5xQZ6FSqgSSHq-m7tNA76-6HvMljXh5yZrGtGDQ8Jv35ZeMKSAdo-kYdjTxOqiUZHtAl43bkInehCVFpL_glZg8l2cmo0n_87sTeoe41oQcizk_XyIeZCwtMD5iOj-8"
}
```

`Request:`

```yml
    "avatar": image binary file
```

`Response:`

```json
{
    "image": "http://35.240.249.239:5000/public/879203e2-ecf2-47ee-a435-96d721f9d852.png"
}
```
