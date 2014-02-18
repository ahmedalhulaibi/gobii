package gobii/gaze

type txUserParam uintptr
type txHandle uintptr
type txTicket int

type txResult int

const (
	txEnumStartValue = 1
	txCleanupTimeDefault = 500
	
	txTrue = 1
	txFalse = 0
	TxEmptyHandle = 0
)

const (
	txResultUnknown txResult = iota + txEnumStartValue
	txResultOk
	txResultSystemNotInitialized
	txResultSystemAlreadyInitialized
	txResultSystemStillInUse
	txResultInvalidArgument
	txResultInvalidContext
	txResultInvalidHandle
	txResultNotFound
	txResultInvalidBufferSize
	txResultDuplicateProperty
	txResultDuplicateBounds
	txResultDuplicateBehaviour
	txResultDuplicateInteractor
	txResultDuplicateSettingsObserver
	txResultInvalidPropertyType
	txResultInvalidPropertyName
	txResultPropertyNotRemovable
	txResultNotConnected
	txResultInvalidObjectCast
	txResultInvalidThread
	txResultInvalidBoundsType
	txResultInvalidBehaviourType
	txResultObjectLeakage
	txResultObjectTrakingNotEnabled
)

var resultString = map[txResult]string{
	txResultUnknown:                   "Unknown",
	txResultOk:                        "Ok",
	txResultSystemNotInitialized:      "System isn't initialized",
	txResultSystemAlreadyInitialized:  "System is already initialized",
	txResultSystemStillInUse:          "System is still in use",
	txResultInvalidArgument:           "Invalid argument",
	txResultInvalidContext:            "Invalid context",
	txResultInvalidHandle:             "Invalid handle",
	txResultNotFound:                  "Not found",
	txResultInvalidBufferSize:         "Invalid buffer size",
	txResultDuplicateProperty:         "Duplicate property",
	txResultDuplicateBounds:           "Duplicate bounds",
	txResultDuplicateBehaviour:        "Duplicate behaviour",
	txResultDuplicateInteractor:       "Duplicate interactor",
	txResultDuplicateSettingsObserver: "Duplicate settings observer",
	txResultInvalidPropertyType:       "Invalid property type",
	txResultInvalidPropertyName:       "Invalid property name",
	txResultPropertyNotRemovable:      "Property not removeable",
	txResultNotConnected:              "Not connected",
	txResultInvalidObjectCast:         "Invalid Object cast",
	txResultInvalidThread:             "Invalid thread",
	txResultInvalidBoundsType:         "Invalid bounds type",
	txResultInvalidBehaviourType:      "Invalid behaviour type",
	txResultObjectLeakage:             "Object leakage",
	txResultObjectTrakingNotEnabled:   "Object tracking not enabled",
}

func (t txResult) Error() string {
	if t < txResultOk || t > txResultObjectTrakingNotEnabled {
		return "Unknown"
	}

	return resultString[t]
}

const (
	txSystemComponentOverrideFlagNone        = 0
	txSystemComponentOverrideFlagMemoryModel = 1 << iota
	txSystemComponentOverrideFlagThreadingModel
	txSystemComponentOverrideFlagLoggingModel
)