namespace go filtergen

enum IDType {
    origin_id = 1,
    md5_id = 2,
    sha1_id = 3,

}

struct AntiSpamRequest {
    1: string id,
    2: IDType idType,   #设备id加密类型
    3: string ip,
    4: list<string> slotIds,
    5: string ua,
    6: string sourceId,
    7: string geo,
    8: i32 devType,
    9: i32 idDt,
}

struct AntiSpamResponse {
    1:list<bool> legals,
    2:list<string> reasons,
}

service IllegalService {
        AntiSpamResponse getIllegalReason(1: AntiSpamRequest request),
}


