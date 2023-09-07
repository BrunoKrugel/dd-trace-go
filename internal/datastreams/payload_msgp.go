package datastreams

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Backlog) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Tags":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Tags")
				return
			}
			if cap(z.Tags) >= int(zb0002) {
				z.Tags = (z.Tags)[:zb0002]
			} else {
				z.Tags = make([]string, zb0002)
			}
			for za0001 := range z.Tags {
				z.Tags[za0001], err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "Tags", za0001)
					return
				}
			}
		case "Value":
			z.Value, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "Value")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Backlog) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Tags"
	err = en.Append(0x82, 0xa4, 0x54, 0x61, 0x67, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Tags)))
	if err != nil {
		err = msgp.WrapError(err, "Tags")
		return
	}
	for za0001 := range z.Tags {
		err = en.WriteString(z.Tags[za0001])
		if err != nil {
			err = msgp.WrapError(err, "Tags", za0001)
			return
		}
	}
	// write "Value"
	err = en.Append(0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Value)
	if err != nil {
		err = msgp.WrapError(err, "Value")
		return
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Backlog) Msgsize() (s int) {
	s = 1 + 5 + msgp.ArrayHeaderSize
	for za0001 := range z.Tags {
		s += msgp.StringPrefixSize + len(z.Tags[za0001])
	}
	s += 6 + msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *CommitOffset) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "ConsumerGroup":
			z.ConsumerGroup, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "ConsumerGroup")
				return
			}
		case "Topic":
			z.Topic, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Topic")
				return
			}
		case "Partition":
			z.Partition, err = dc.ReadInt32()
			if err != nil {
				err = msgp.WrapError(err, "Partition")
				return
			}
		case "Offset":
			z.Offset, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "Offset")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *CommitOffset) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "ConsumerGroup"
	err = en.Append(0x84, 0xad, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70)
	if err != nil {
		return
	}
	err = en.WriteString(z.ConsumerGroup)
	if err != nil {
		err = msgp.WrapError(err, "ConsumerGroup")
		return
	}
	// write "Topic"
	err = en.Append(0xa5, 0x54, 0x6f, 0x70, 0x69, 0x63)
	if err != nil {
		return
	}
	err = en.WriteString(z.Topic)
	if err != nil {
		err = msgp.WrapError(err, "Topic")
		return
	}
	// write "Partition"
	err = en.Append(0xa9, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e)
	if err != nil {
		return
	}
	err = en.WriteInt32(z.Partition)
	if err != nil {
		err = msgp.WrapError(err, "Partition")
		return
	}
	// write "Offset"
	err = en.Append(0xa6, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Offset)
	if err != nil {
		err = msgp.WrapError(err, "Offset")
		return
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *CommitOffset) Msgsize() (s int) {
	s = 1 + 14 + msgp.StringPrefixSize + len(z.ConsumerGroup) + 6 + msgp.StringPrefixSize + len(z.Topic) + 10 + msgp.Int32Size + 7 + msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ProduceOffset) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Topic":
			z.Topic, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Topic")
				return
			}
		case "Partition":
			z.Partition, err = dc.ReadInt32()
			if err != nil {
				err = msgp.WrapError(err, "Partition")
				return
			}
		case "Offset":
			z.Offset, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "Offset")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z ProduceOffset) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "Topic"
	err = en.Append(0x83, 0xa5, 0x54, 0x6f, 0x70, 0x69, 0x63)
	if err != nil {
		return
	}
	err = en.WriteString(z.Topic)
	if err != nil {
		err = msgp.WrapError(err, "Topic")
		return
	}
	// write "Partition"
	err = en.Append(0xa9, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e)
	if err != nil {
		return
	}
	err = en.WriteInt32(z.Partition)
	if err != nil {
		err = msgp.WrapError(err, "Partition")
		return
	}
	// write "Offset"
	err = en.Append(0xa6, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Offset)
	if err != nil {
		err = msgp.WrapError(err, "Offset")
		return
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z ProduceOffset) Msgsize() (s int) {
	s = 1 + 6 + msgp.StringPrefixSize + len(z.Topic) + 10 + msgp.Int32Size + 7 + msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *StatsBucket) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Start":
			z.Start, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "Start")
				return
			}
		case "Duration":
			z.Duration, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "Duration")
				return
			}
		case "Stats":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Stats")
				return
			}
			if cap(z.Stats) >= int(zb0002) {
				z.Stats = (z.Stats)[:zb0002]
			} else {
				z.Stats = make([]StatsPoint, zb0002)
			}
			for za0001 := range z.Stats {
				err = z.Stats[za0001].DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Stats", za0001)
					return
				}
			}
		case "Backlogs":
			var zb0003 uint32
			zb0003, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Backlogs")
				return
			}
			if cap(z.Backlogs) >= int(zb0003) {
				z.Backlogs = (z.Backlogs)[:zb0003]
			} else {
				z.Backlogs = make([]Backlog, zb0003)
			}
			for za0002 := range z.Backlogs {
				var zb0004 uint32
				zb0004, err = dc.ReadMapHeader()
				if err != nil {
					err = msgp.WrapError(err, "Backlogs", za0002)
					return
				}
				for zb0004 > 0 {
					zb0004--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						err = msgp.WrapError(err, "Backlogs", za0002)
						return
					}
					switch msgp.UnsafeString(field) {
					case "Tags":
						var zb0005 uint32
						zb0005, err = dc.ReadArrayHeader()
						if err != nil {
							err = msgp.WrapError(err, "Backlogs", za0002, "Tags")
							return
						}
						if cap(z.Backlogs[za0002].Tags) >= int(zb0005) {
							z.Backlogs[za0002].Tags = (z.Backlogs[za0002].Tags)[:zb0005]
						} else {
							z.Backlogs[za0002].Tags = make([]string, zb0005)
						}
						for za0003 := range z.Backlogs[za0002].Tags {
							z.Backlogs[za0002].Tags[za0003], err = dc.ReadString()
							if err != nil {
								err = msgp.WrapError(err, "Backlogs", za0002, "Tags", za0003)
								return
							}
						}
					case "Value":
						z.Backlogs[za0002].Value, err = dc.ReadInt64()
						if err != nil {
							err = msgp.WrapError(err, "Backlogs", za0002, "Value")
							return
						}
					default:
						err = dc.Skip()
						if err != nil {
							err = msgp.WrapError(err, "Backlogs", za0002)
							return
						}
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *StatsBucket) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "Start"
	err = en.Append(0x84, 0xa5, 0x53, 0x74, 0x61, 0x72, 0x74)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Start)
	if err != nil {
		err = msgp.WrapError(err, "Start")
		return
	}
	// write "Duration"
	err = en.Append(0xa8, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Duration)
	if err != nil {
		err = msgp.WrapError(err, "Duration")
		return
	}
	// write "Stats"
	err = en.Append(0xa5, 0x53, 0x74, 0x61, 0x74, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Stats)))
	if err != nil {
		err = msgp.WrapError(err, "Stats")
		return
	}
	for za0001 := range z.Stats {
		err = z.Stats[za0001].EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Stats", za0001)
			return
		}
	}
	// write "Backlogs"
	err = en.Append(0xa8, 0x42, 0x61, 0x63, 0x6b, 0x6c, 0x6f, 0x67, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Backlogs)))
	if err != nil {
		err = msgp.WrapError(err, "Backlogs")
		return
	}
	for za0002 := range z.Backlogs {
		// map header, size 2
		// write "Tags"
		err = en.Append(0x82, 0xa4, 0x54, 0x61, 0x67, 0x73)
		if err != nil {
			return
		}
		err = en.WriteArrayHeader(uint32(len(z.Backlogs[za0002].Tags)))
		if err != nil {
			err = msgp.WrapError(err, "Backlogs", za0002, "Tags")
			return
		}
		for za0003 := range z.Backlogs[za0002].Tags {
			err = en.WriteString(z.Backlogs[za0002].Tags[za0003])
			if err != nil {
				err = msgp.WrapError(err, "Backlogs", za0002, "Tags", za0003)
				return
			}
		}
		// write "Value"
		err = en.Append(0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
		if err != nil {
			return
		}
		err = en.WriteInt64(z.Backlogs[za0002].Value)
		if err != nil {
			err = msgp.WrapError(err, "Backlogs", za0002, "Value")
			return
		}
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *StatsBucket) Msgsize() (s int) {
	s = 1 + 6 + msgp.Uint64Size + 9 + msgp.Uint64Size + 6 + msgp.ArrayHeaderSize
	for za0001 := range z.Stats {
		s += z.Stats[za0001].Msgsize()
	}
	s += 9 + msgp.ArrayHeaderSize
	for za0002 := range z.Backlogs {
		s += 1 + 5 + msgp.ArrayHeaderSize
		for za0003 := range z.Backlogs[za0002].Tags {
			s += msgp.StringPrefixSize + len(z.Backlogs[za0002].Tags[za0003])
		}
		s += 6 + msgp.Int64Size
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *StatsPayload) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Env":
			z.Env, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Env")
				return
			}
		case "Service":
			z.Service, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Service")
				return
			}
		case "Stats":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Stats")
				return
			}
			if cap(z.Stats) >= int(zb0002) {
				z.Stats = (z.Stats)[:zb0002]
			} else {
				z.Stats = make([]StatsBucket, zb0002)
			}
			for za0001 := range z.Stats {
				err = z.Stats[za0001].DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Stats", za0001)
					return
				}
			}
		case "TracerVersion":
			z.TracerVersion, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "TracerVersion")
				return
			}
		case "Lang":
			z.Lang, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Lang")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *StatsPayload) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 5
	// write "Env"
	err = en.Append(0x85, 0xa3, 0x45, 0x6e, 0x76)
	if err != nil {
		return
	}
	err = en.WriteString(z.Env)
	if err != nil {
		err = msgp.WrapError(err, "Env")
		return
	}
	// write "Service"
	err = en.Append(0xa7, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Service)
	if err != nil {
		err = msgp.WrapError(err, "Service")
		return
	}
	// write "Stats"
	err = en.Append(0xa5, 0x53, 0x74, 0x61, 0x74, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Stats)))
	if err != nil {
		err = msgp.WrapError(err, "Stats")
		return
	}
	for za0001 := range z.Stats {
		err = z.Stats[za0001].EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Stats", za0001)
			return
		}
	}
	// write "TracerVersion"
	err = en.Append(0xad, 0x54, 0x72, 0x61, 0x63, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
	if err != nil {
		return
	}
	err = en.WriteString(z.TracerVersion)
	if err != nil {
		err = msgp.WrapError(err, "TracerVersion")
		return
	}
	// write "Lang"
	err = en.Append(0xa4, 0x4c, 0x61, 0x6e, 0x67)
	if err != nil {
		return
	}
	err = en.WriteString(z.Lang)
	if err != nil {
		err = msgp.WrapError(err, "Lang")
		return
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *StatsPayload) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(z.Env) + 8 + msgp.StringPrefixSize + len(z.Service) + 6 + msgp.ArrayHeaderSize
	for za0001 := range z.Stats {
		s += z.Stats[za0001].Msgsize()
	}
	s += 14 + msgp.StringPrefixSize + len(z.TracerVersion) + 5 + msgp.StringPrefixSize + len(z.Lang)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *StatsPoint) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Service":
			z.Service, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Service")
				return
			}
		case "EdgeTags":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "EdgeTags")
				return
			}
			if cap(z.EdgeTags) >= int(zb0002) {
				z.EdgeTags = (z.EdgeTags)[:zb0002]
			} else {
				z.EdgeTags = make([]string, zb0002)
			}
			for za0001 := range z.EdgeTags {
				z.EdgeTags[za0001], err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "EdgeTags", za0001)
					return
				}
			}
		case "Hash":
			z.Hash, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "Hash")
				return
			}
		case "ParentHash":
			z.ParentHash, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "ParentHash")
				return
			}
		case "PathwayLatency":
			z.PathwayLatency, err = dc.ReadBytes(z.PathwayLatency)
			if err != nil {
				err = msgp.WrapError(err, "PathwayLatency")
				return
			}
		case "EdgeLatency":
			z.EdgeLatency, err = dc.ReadBytes(z.EdgeLatency)
			if err != nil {
				err = msgp.WrapError(err, "EdgeLatency")
				return
			}
		case "PayloadSize":
			z.PayloadSize, err = dc.ReadBytes(z.PayloadSize)
			if err != nil {
				err = msgp.WrapError(err, "PayloadSize")
				return
			}
		case "TimestampType":
			{
				var zb0003 string
				zb0003, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "TimestampType")
					return
				}
				z.TimestampType = TimestampType(zb0003)
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *StatsPoint) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 8
	// write "Service"
	err = en.Append(0x88, 0xa7, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Service)
	if err != nil {
		err = msgp.WrapError(err, "Service")
		return
	}
	// write "EdgeTags"
	err = en.Append(0xa8, 0x45, 0x64, 0x67, 0x65, 0x54, 0x61, 0x67, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.EdgeTags)))
	if err != nil {
		err = msgp.WrapError(err, "EdgeTags")
		return
	}
	for za0001 := range z.EdgeTags {
		err = en.WriteString(z.EdgeTags[za0001])
		if err != nil {
			err = msgp.WrapError(err, "EdgeTags", za0001)
			return
		}
	}
	// write "Hash"
	err = en.Append(0xa4, 0x48, 0x61, 0x73, 0x68)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Hash)
	if err != nil {
		err = msgp.WrapError(err, "Hash")
		return
	}
	// write "ParentHash"
	err = en.Append(0xaa, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.ParentHash)
	if err != nil {
		err = msgp.WrapError(err, "ParentHash")
		return
	}
	// write "PathwayLatency"
	err = en.Append(0xae, 0x50, 0x61, 0x74, 0x68, 0x77, 0x61, 0x79, 0x4c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.PathwayLatency)
	if err != nil {
		err = msgp.WrapError(err, "PathwayLatency")
		return
	}
	// write "EdgeLatency"
	err = en.Append(0xab, 0x45, 0x64, 0x67, 0x65, 0x4c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.EdgeLatency)
	if err != nil {
		err = msgp.WrapError(err, "EdgeLatency")
		return
	}
	// write "PayloadSize"
	err = en.Append(0xab, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x69, 0x7a, 0x65)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.PayloadSize)
	if err != nil {
		err = msgp.WrapError(err, "PayloadSize")
		return
	}
	// write "TimestampType"
	err = en.Append(0xad, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x54, 0x79, 0x70, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(string(z.TimestampType))
	if err != nil {
		err = msgp.WrapError(err, "TimestampType")
		return
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *StatsPoint) Msgsize() (s int) {
	s = 1 + 8 + msgp.StringPrefixSize + len(z.Service) + 9 + msgp.ArrayHeaderSize
	for za0001 := range z.EdgeTags {
		s += msgp.StringPrefixSize + len(z.EdgeTags[za0001])
	}
	s += 5 + msgp.Uint64Size + 11 + msgp.Uint64Size + 15 + msgp.BytesPrefixSize + len(z.PathwayLatency) + 12 + msgp.BytesPrefixSize + len(z.EdgeLatency) + 12 + msgp.BytesPrefixSize + len(z.PayloadSize) + 14 + msgp.StringPrefixSize + len(string(z.TimestampType))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *TimestampType) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 string
		zb0001, err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = TimestampType(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z TimestampType) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteString(string(z))
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z TimestampType) Msgsize() (s int) {
	s = msgp.StringPrefixSize + len(string(z))
	return
}