# Swapro Technical Test
---
## Fitur
### Employee (Karyawan)
- **Register**: Mendaftarkan karyawan baru ke dalam sistem.
- **AssignSuperior**: Menetapkan atasan untuk karyawan.
- **GetInformation**: Mendapatkan informasi detail mengenai karyawan.
- **Delete**: Menghapus data karyawan dari sistem.

### Position (Posisi)
- **GetPositionInformation**: Mendapatkan informasi detail tentang posisi tertentu.
- **ApplyToPositionInDepartment**: Mengajukan posisi dalam suatu departemen.
- **ChangePositionName**: Mengubah nama posisi yang ada.
- **DeletePosition**: Menghapus posisi dari sistem.

### Department (Departemen)
- **GetDepartmentInformation**: Mendapatkan informasi detail tentang departemen tertentu.
- **ChangeDepartmentName**: Mengubah nama departemen yang ada.
- **DeleteDepartment**: Menghapus departemen dari sistem.

### Attendance (Kehadiran)
- **CheckInToLocation**: Mencatat kehadiran karyawan di lokasi tertentu.
- **CheckOut**: Mencatat waktu keluar karyawan dari lokasi.
- **GetReportOfAttendances**: Mendapatkan laporan kehadiran karyawan.
- **DeleteAttendances**: Menghapus data kehadiran karyawan.

### Location (Lokasi)
- **GetLocationsEverAttended**: Mendapatkan daftar lokasi yang pernah dikunjungi oleh karyawan.
- **ChangeLocationNameByAttendance**: Mengubah nama lokasi berdasarkan data kehadiran.
- **DeleteLocationByAttendance**: Menghapus data lokasi berdasarkan kehadiran karyawan.

## Skema Database

![Database Schema](public/schema.png)

## Catatan

Table `employee` tidak perlu ada `department id`. Karena sudah memiliki `position id`, yang dimana table `position` sudah memiliki `department id`. Jika tetap ada, ada kemungkinan data terduplikasi. Sebagai contoh:

- Employee dengan id 1 memiliki department dengan id 3 dan position id 5
- Sedangkan position id 5 memiliki department id 9

Lalu pertanyaannya, employee bekerja di department id 3 atau 9?
