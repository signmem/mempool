package db

import "log"

func QueryHosts() (map[string]int64, error) {
	m := make(map[string]int64)
	sql := "select hostname, UNIX_TIMESTAMP() timestamp from host"
	rows, err := DB.Query(sql)
	if err != nil {
		log.Printf("[ERROR] QueryHosts() dbquery err:%s\n", err)
	}

	defer rows.Close()
	for rows.Next() {
		var (
			hostname 		string
			timestamp 		int64
		)

		err = rows.Scan(&hostname, &timestamp)
		if err != nil {
			log.Printf("[ERROR] QueryHosts() db scan err:%s\n", err)
			continue
		}

		m[hostname] = timestamp
	}
	return m, nil
}