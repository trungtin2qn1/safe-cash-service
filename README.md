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
        "id": "1620868265",
        "content": "Vừa có người dùng muốn mở khóa",
        "method": "face",
        "is_success": true,
        "user_id": "1620448417",
        "created_at": "2020-07-30T18:22:54.42639Z",
        "updated_at": "2020-07-30T18:22:54.42639Z",
        "first_name": "Tin",
        "last_name": "Huynh"
    }
]
```

**16. Watch file in browser**

`URL`: `/public/{filename}`
