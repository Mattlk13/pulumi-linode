// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.linode.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.linode.outputs.GetDatabaseMysqlConfigMysqlConnectTimeout;
import com.pulumi.linode.outputs.GetDatabaseMysqlConfigMysqlDefaultTimeZone;
import com.pulumi.linode.outputs.GetDatabaseMysqlConfigMysqlGroupConcatMaxLen;
import com.pulumi.linode.outputs.GetDatabaseMysqlConfigMysqlInformationSchemaStatsExpiry;
import com.pulumi.linode.outputs.GetDatabaseMysqlConfigMysqlInnodbChangeBufferMaxSize;
import com.pulumi.linode.outputs.GetDatabaseMysqlConfigMysqlInnodbFlushNeighbors;
import com.pulumi.linode.outputs.GetDatabaseMysqlConfigMysqlInnodbFtMinTokenSize;
import com.pulumi.linode.outputs.GetDatabaseMysqlConfigMysqlInnodbFtServerStopwordTable;
import com.pulumi.linode.outputs.GetDatabaseMysqlConfigMysqlInnodbLockWaitTimeout;
import com.pulumi.linode.outputs.GetDatabaseMysqlConfigMysqlInnodbLogBufferSize;
import com.pulumi.linode.outputs.GetDatabaseMysqlConfigMysqlInnodbOnlineAlterLogMaxSize;
import com.pulumi.linode.outputs.GetDatabaseMysqlConfigMysqlInnodbReadIoThreads;
import com.pulumi.linode.outputs.GetDatabaseMysqlConfigMysqlInnodbRollbackOnTimeout;
import com.pulumi.linode.outputs.GetDatabaseMysqlConfigMysqlInnodbThreadConcurrency;
import com.pulumi.linode.outputs.GetDatabaseMysqlConfigMysqlInnodbWriteIoThreads;
import com.pulumi.linode.outputs.GetDatabaseMysqlConfigMysqlInteractiveTimeout;
import com.pulumi.linode.outputs.GetDatabaseMysqlConfigMysqlInternalTmpMemStorageEngine;
import com.pulumi.linode.outputs.GetDatabaseMysqlConfigMysqlMaxAllowedPacket;
import com.pulumi.linode.outputs.GetDatabaseMysqlConfigMysqlMaxHeapTableSize;
import com.pulumi.linode.outputs.GetDatabaseMysqlConfigMysqlNetBufferLength;
import com.pulumi.linode.outputs.GetDatabaseMysqlConfigMysqlNetReadTimeout;
import com.pulumi.linode.outputs.GetDatabaseMysqlConfigMysqlNetWriteTimeout;
import com.pulumi.linode.outputs.GetDatabaseMysqlConfigMysqlSortBufferSize;
import com.pulumi.linode.outputs.GetDatabaseMysqlConfigMysqlSqlMode;
import com.pulumi.linode.outputs.GetDatabaseMysqlConfigMysqlSqlRequirePrimaryKey;
import com.pulumi.linode.outputs.GetDatabaseMysqlConfigMysqlTmpTableSize;
import com.pulumi.linode.outputs.GetDatabaseMysqlConfigMysqlWaitTimeout;
import java.util.Objects;

@CustomType
public final class GetDatabaseMysqlConfigMysql {
    private GetDatabaseMysqlConfigMysqlConnectTimeout connectTimeout;
    private GetDatabaseMysqlConfigMysqlDefaultTimeZone defaultTimeZone;
    private GetDatabaseMysqlConfigMysqlGroupConcatMaxLen groupConcatMaxLen;
    private GetDatabaseMysqlConfigMysqlInformationSchemaStatsExpiry informationSchemaStatsExpiry;
    private GetDatabaseMysqlConfigMysqlInnodbChangeBufferMaxSize innodbChangeBufferMaxSize;
    private GetDatabaseMysqlConfigMysqlInnodbFlushNeighbors innodbFlushNeighbors;
    private GetDatabaseMysqlConfigMysqlInnodbFtMinTokenSize innodbFtMinTokenSize;
    private GetDatabaseMysqlConfigMysqlInnodbFtServerStopwordTable innodbFtServerStopwordTable;
    private GetDatabaseMysqlConfigMysqlInnodbLockWaitTimeout innodbLockWaitTimeout;
    private GetDatabaseMysqlConfigMysqlInnodbLogBufferSize innodbLogBufferSize;
    private GetDatabaseMysqlConfigMysqlInnodbOnlineAlterLogMaxSize innodbOnlineAlterLogMaxSize;
    private GetDatabaseMysqlConfigMysqlInnodbReadIoThreads innodbReadIoThreads;
    private GetDatabaseMysqlConfigMysqlInnodbRollbackOnTimeout innodbRollbackOnTimeout;
    private GetDatabaseMysqlConfigMysqlInnodbThreadConcurrency innodbThreadConcurrency;
    private GetDatabaseMysqlConfigMysqlInnodbWriteIoThreads innodbWriteIoThreads;
    private GetDatabaseMysqlConfigMysqlInteractiveTimeout interactiveTimeout;
    private GetDatabaseMysqlConfigMysqlInternalTmpMemStorageEngine internalTmpMemStorageEngine;
    private GetDatabaseMysqlConfigMysqlMaxAllowedPacket maxAllowedPacket;
    private GetDatabaseMysqlConfigMysqlMaxHeapTableSize maxHeapTableSize;
    private GetDatabaseMysqlConfigMysqlNetBufferLength netBufferLength;
    private GetDatabaseMysqlConfigMysqlNetReadTimeout netReadTimeout;
    private GetDatabaseMysqlConfigMysqlNetWriteTimeout netWriteTimeout;
    private GetDatabaseMysqlConfigMysqlSortBufferSize sortBufferSize;
    private GetDatabaseMysqlConfigMysqlSqlMode sqlMode;
    private GetDatabaseMysqlConfigMysqlSqlRequirePrimaryKey sqlRequirePrimaryKey;
    private GetDatabaseMysqlConfigMysqlTmpTableSize tmpTableSize;
    private GetDatabaseMysqlConfigMysqlWaitTimeout waitTimeout;

    private GetDatabaseMysqlConfigMysql() {}
    public GetDatabaseMysqlConfigMysqlConnectTimeout connectTimeout() {
        return this.connectTimeout;
    }
    public GetDatabaseMysqlConfigMysqlDefaultTimeZone defaultTimeZone() {
        return this.defaultTimeZone;
    }
    public GetDatabaseMysqlConfigMysqlGroupConcatMaxLen groupConcatMaxLen() {
        return this.groupConcatMaxLen;
    }
    public GetDatabaseMysqlConfigMysqlInformationSchemaStatsExpiry informationSchemaStatsExpiry() {
        return this.informationSchemaStatsExpiry;
    }
    public GetDatabaseMysqlConfigMysqlInnodbChangeBufferMaxSize innodbChangeBufferMaxSize() {
        return this.innodbChangeBufferMaxSize;
    }
    public GetDatabaseMysqlConfigMysqlInnodbFlushNeighbors innodbFlushNeighbors() {
        return this.innodbFlushNeighbors;
    }
    public GetDatabaseMysqlConfigMysqlInnodbFtMinTokenSize innodbFtMinTokenSize() {
        return this.innodbFtMinTokenSize;
    }
    public GetDatabaseMysqlConfigMysqlInnodbFtServerStopwordTable innodbFtServerStopwordTable() {
        return this.innodbFtServerStopwordTable;
    }
    public GetDatabaseMysqlConfigMysqlInnodbLockWaitTimeout innodbLockWaitTimeout() {
        return this.innodbLockWaitTimeout;
    }
    public GetDatabaseMysqlConfigMysqlInnodbLogBufferSize innodbLogBufferSize() {
        return this.innodbLogBufferSize;
    }
    public GetDatabaseMysqlConfigMysqlInnodbOnlineAlterLogMaxSize innodbOnlineAlterLogMaxSize() {
        return this.innodbOnlineAlterLogMaxSize;
    }
    public GetDatabaseMysqlConfigMysqlInnodbReadIoThreads innodbReadIoThreads() {
        return this.innodbReadIoThreads;
    }
    public GetDatabaseMysqlConfigMysqlInnodbRollbackOnTimeout innodbRollbackOnTimeout() {
        return this.innodbRollbackOnTimeout;
    }
    public GetDatabaseMysqlConfigMysqlInnodbThreadConcurrency innodbThreadConcurrency() {
        return this.innodbThreadConcurrency;
    }
    public GetDatabaseMysqlConfigMysqlInnodbWriteIoThreads innodbWriteIoThreads() {
        return this.innodbWriteIoThreads;
    }
    public GetDatabaseMysqlConfigMysqlInteractiveTimeout interactiveTimeout() {
        return this.interactiveTimeout;
    }
    public GetDatabaseMysqlConfigMysqlInternalTmpMemStorageEngine internalTmpMemStorageEngine() {
        return this.internalTmpMemStorageEngine;
    }
    public GetDatabaseMysqlConfigMysqlMaxAllowedPacket maxAllowedPacket() {
        return this.maxAllowedPacket;
    }
    public GetDatabaseMysqlConfigMysqlMaxHeapTableSize maxHeapTableSize() {
        return this.maxHeapTableSize;
    }
    public GetDatabaseMysqlConfigMysqlNetBufferLength netBufferLength() {
        return this.netBufferLength;
    }
    public GetDatabaseMysqlConfigMysqlNetReadTimeout netReadTimeout() {
        return this.netReadTimeout;
    }
    public GetDatabaseMysqlConfigMysqlNetWriteTimeout netWriteTimeout() {
        return this.netWriteTimeout;
    }
    public GetDatabaseMysqlConfigMysqlSortBufferSize sortBufferSize() {
        return this.sortBufferSize;
    }
    public GetDatabaseMysqlConfigMysqlSqlMode sqlMode() {
        return this.sqlMode;
    }
    public GetDatabaseMysqlConfigMysqlSqlRequirePrimaryKey sqlRequirePrimaryKey() {
        return this.sqlRequirePrimaryKey;
    }
    public GetDatabaseMysqlConfigMysqlTmpTableSize tmpTableSize() {
        return this.tmpTableSize;
    }
    public GetDatabaseMysqlConfigMysqlWaitTimeout waitTimeout() {
        return this.waitTimeout;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDatabaseMysqlConfigMysql defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private GetDatabaseMysqlConfigMysqlConnectTimeout connectTimeout;
        private GetDatabaseMysqlConfigMysqlDefaultTimeZone defaultTimeZone;
        private GetDatabaseMysqlConfigMysqlGroupConcatMaxLen groupConcatMaxLen;
        private GetDatabaseMysqlConfigMysqlInformationSchemaStatsExpiry informationSchemaStatsExpiry;
        private GetDatabaseMysqlConfigMysqlInnodbChangeBufferMaxSize innodbChangeBufferMaxSize;
        private GetDatabaseMysqlConfigMysqlInnodbFlushNeighbors innodbFlushNeighbors;
        private GetDatabaseMysqlConfigMysqlInnodbFtMinTokenSize innodbFtMinTokenSize;
        private GetDatabaseMysqlConfigMysqlInnodbFtServerStopwordTable innodbFtServerStopwordTable;
        private GetDatabaseMysqlConfigMysqlInnodbLockWaitTimeout innodbLockWaitTimeout;
        private GetDatabaseMysqlConfigMysqlInnodbLogBufferSize innodbLogBufferSize;
        private GetDatabaseMysqlConfigMysqlInnodbOnlineAlterLogMaxSize innodbOnlineAlterLogMaxSize;
        private GetDatabaseMysqlConfigMysqlInnodbReadIoThreads innodbReadIoThreads;
        private GetDatabaseMysqlConfigMysqlInnodbRollbackOnTimeout innodbRollbackOnTimeout;
        private GetDatabaseMysqlConfigMysqlInnodbThreadConcurrency innodbThreadConcurrency;
        private GetDatabaseMysqlConfigMysqlInnodbWriteIoThreads innodbWriteIoThreads;
        private GetDatabaseMysqlConfigMysqlInteractiveTimeout interactiveTimeout;
        private GetDatabaseMysqlConfigMysqlInternalTmpMemStorageEngine internalTmpMemStorageEngine;
        private GetDatabaseMysqlConfigMysqlMaxAllowedPacket maxAllowedPacket;
        private GetDatabaseMysqlConfigMysqlMaxHeapTableSize maxHeapTableSize;
        private GetDatabaseMysqlConfigMysqlNetBufferLength netBufferLength;
        private GetDatabaseMysqlConfigMysqlNetReadTimeout netReadTimeout;
        private GetDatabaseMysqlConfigMysqlNetWriteTimeout netWriteTimeout;
        private GetDatabaseMysqlConfigMysqlSortBufferSize sortBufferSize;
        private GetDatabaseMysqlConfigMysqlSqlMode sqlMode;
        private GetDatabaseMysqlConfigMysqlSqlRequirePrimaryKey sqlRequirePrimaryKey;
        private GetDatabaseMysqlConfigMysqlTmpTableSize tmpTableSize;
        private GetDatabaseMysqlConfigMysqlWaitTimeout waitTimeout;
        public Builder() {}
        public Builder(GetDatabaseMysqlConfigMysql defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.connectTimeout = defaults.connectTimeout;
    	      this.defaultTimeZone = defaults.defaultTimeZone;
    	      this.groupConcatMaxLen = defaults.groupConcatMaxLen;
    	      this.informationSchemaStatsExpiry = defaults.informationSchemaStatsExpiry;
    	      this.innodbChangeBufferMaxSize = defaults.innodbChangeBufferMaxSize;
    	      this.innodbFlushNeighbors = defaults.innodbFlushNeighbors;
    	      this.innodbFtMinTokenSize = defaults.innodbFtMinTokenSize;
    	      this.innodbFtServerStopwordTable = defaults.innodbFtServerStopwordTable;
    	      this.innodbLockWaitTimeout = defaults.innodbLockWaitTimeout;
    	      this.innodbLogBufferSize = defaults.innodbLogBufferSize;
    	      this.innodbOnlineAlterLogMaxSize = defaults.innodbOnlineAlterLogMaxSize;
    	      this.innodbReadIoThreads = defaults.innodbReadIoThreads;
    	      this.innodbRollbackOnTimeout = defaults.innodbRollbackOnTimeout;
    	      this.innodbThreadConcurrency = defaults.innodbThreadConcurrency;
    	      this.innodbWriteIoThreads = defaults.innodbWriteIoThreads;
    	      this.interactiveTimeout = defaults.interactiveTimeout;
    	      this.internalTmpMemStorageEngine = defaults.internalTmpMemStorageEngine;
    	      this.maxAllowedPacket = defaults.maxAllowedPacket;
    	      this.maxHeapTableSize = defaults.maxHeapTableSize;
    	      this.netBufferLength = defaults.netBufferLength;
    	      this.netReadTimeout = defaults.netReadTimeout;
    	      this.netWriteTimeout = defaults.netWriteTimeout;
    	      this.sortBufferSize = defaults.sortBufferSize;
    	      this.sqlMode = defaults.sqlMode;
    	      this.sqlRequirePrimaryKey = defaults.sqlRequirePrimaryKey;
    	      this.tmpTableSize = defaults.tmpTableSize;
    	      this.waitTimeout = defaults.waitTimeout;
        }

        @CustomType.Setter
        public Builder connectTimeout(GetDatabaseMysqlConfigMysqlConnectTimeout connectTimeout) {
            if (connectTimeout == null) {
              throw new MissingRequiredPropertyException("GetDatabaseMysqlConfigMysql", "connectTimeout");
            }
            this.connectTimeout = connectTimeout;
            return this;
        }
        @CustomType.Setter
        public Builder defaultTimeZone(GetDatabaseMysqlConfigMysqlDefaultTimeZone defaultTimeZone) {
            if (defaultTimeZone == null) {
              throw new MissingRequiredPropertyException("GetDatabaseMysqlConfigMysql", "defaultTimeZone");
            }
            this.defaultTimeZone = defaultTimeZone;
            return this;
        }
        @CustomType.Setter
        public Builder groupConcatMaxLen(GetDatabaseMysqlConfigMysqlGroupConcatMaxLen groupConcatMaxLen) {
            if (groupConcatMaxLen == null) {
              throw new MissingRequiredPropertyException("GetDatabaseMysqlConfigMysql", "groupConcatMaxLen");
            }
            this.groupConcatMaxLen = groupConcatMaxLen;
            return this;
        }
        @CustomType.Setter
        public Builder informationSchemaStatsExpiry(GetDatabaseMysqlConfigMysqlInformationSchemaStatsExpiry informationSchemaStatsExpiry) {
            if (informationSchemaStatsExpiry == null) {
              throw new MissingRequiredPropertyException("GetDatabaseMysqlConfigMysql", "informationSchemaStatsExpiry");
            }
            this.informationSchemaStatsExpiry = informationSchemaStatsExpiry;
            return this;
        }
        @CustomType.Setter
        public Builder innodbChangeBufferMaxSize(GetDatabaseMysqlConfigMysqlInnodbChangeBufferMaxSize innodbChangeBufferMaxSize) {
            if (innodbChangeBufferMaxSize == null) {
              throw new MissingRequiredPropertyException("GetDatabaseMysqlConfigMysql", "innodbChangeBufferMaxSize");
            }
            this.innodbChangeBufferMaxSize = innodbChangeBufferMaxSize;
            return this;
        }
        @CustomType.Setter
        public Builder innodbFlushNeighbors(GetDatabaseMysqlConfigMysqlInnodbFlushNeighbors innodbFlushNeighbors) {
            if (innodbFlushNeighbors == null) {
              throw new MissingRequiredPropertyException("GetDatabaseMysqlConfigMysql", "innodbFlushNeighbors");
            }
            this.innodbFlushNeighbors = innodbFlushNeighbors;
            return this;
        }
        @CustomType.Setter
        public Builder innodbFtMinTokenSize(GetDatabaseMysqlConfigMysqlInnodbFtMinTokenSize innodbFtMinTokenSize) {
            if (innodbFtMinTokenSize == null) {
              throw new MissingRequiredPropertyException("GetDatabaseMysqlConfigMysql", "innodbFtMinTokenSize");
            }
            this.innodbFtMinTokenSize = innodbFtMinTokenSize;
            return this;
        }
        @CustomType.Setter
        public Builder innodbFtServerStopwordTable(GetDatabaseMysqlConfigMysqlInnodbFtServerStopwordTable innodbFtServerStopwordTable) {
            if (innodbFtServerStopwordTable == null) {
              throw new MissingRequiredPropertyException("GetDatabaseMysqlConfigMysql", "innodbFtServerStopwordTable");
            }
            this.innodbFtServerStopwordTable = innodbFtServerStopwordTable;
            return this;
        }
        @CustomType.Setter
        public Builder innodbLockWaitTimeout(GetDatabaseMysqlConfigMysqlInnodbLockWaitTimeout innodbLockWaitTimeout) {
            if (innodbLockWaitTimeout == null) {
              throw new MissingRequiredPropertyException("GetDatabaseMysqlConfigMysql", "innodbLockWaitTimeout");
            }
            this.innodbLockWaitTimeout = innodbLockWaitTimeout;
            return this;
        }
        @CustomType.Setter
        public Builder innodbLogBufferSize(GetDatabaseMysqlConfigMysqlInnodbLogBufferSize innodbLogBufferSize) {
            if (innodbLogBufferSize == null) {
              throw new MissingRequiredPropertyException("GetDatabaseMysqlConfigMysql", "innodbLogBufferSize");
            }
            this.innodbLogBufferSize = innodbLogBufferSize;
            return this;
        }
        @CustomType.Setter
        public Builder innodbOnlineAlterLogMaxSize(GetDatabaseMysqlConfigMysqlInnodbOnlineAlterLogMaxSize innodbOnlineAlterLogMaxSize) {
            if (innodbOnlineAlterLogMaxSize == null) {
              throw new MissingRequiredPropertyException("GetDatabaseMysqlConfigMysql", "innodbOnlineAlterLogMaxSize");
            }
            this.innodbOnlineAlterLogMaxSize = innodbOnlineAlterLogMaxSize;
            return this;
        }
        @CustomType.Setter
        public Builder innodbReadIoThreads(GetDatabaseMysqlConfigMysqlInnodbReadIoThreads innodbReadIoThreads) {
            if (innodbReadIoThreads == null) {
              throw new MissingRequiredPropertyException("GetDatabaseMysqlConfigMysql", "innodbReadIoThreads");
            }
            this.innodbReadIoThreads = innodbReadIoThreads;
            return this;
        }
        @CustomType.Setter
        public Builder innodbRollbackOnTimeout(GetDatabaseMysqlConfigMysqlInnodbRollbackOnTimeout innodbRollbackOnTimeout) {
            if (innodbRollbackOnTimeout == null) {
              throw new MissingRequiredPropertyException("GetDatabaseMysqlConfigMysql", "innodbRollbackOnTimeout");
            }
            this.innodbRollbackOnTimeout = innodbRollbackOnTimeout;
            return this;
        }
        @CustomType.Setter
        public Builder innodbThreadConcurrency(GetDatabaseMysqlConfigMysqlInnodbThreadConcurrency innodbThreadConcurrency) {
            if (innodbThreadConcurrency == null) {
              throw new MissingRequiredPropertyException("GetDatabaseMysqlConfigMysql", "innodbThreadConcurrency");
            }
            this.innodbThreadConcurrency = innodbThreadConcurrency;
            return this;
        }
        @CustomType.Setter
        public Builder innodbWriteIoThreads(GetDatabaseMysqlConfigMysqlInnodbWriteIoThreads innodbWriteIoThreads) {
            if (innodbWriteIoThreads == null) {
              throw new MissingRequiredPropertyException("GetDatabaseMysqlConfigMysql", "innodbWriteIoThreads");
            }
            this.innodbWriteIoThreads = innodbWriteIoThreads;
            return this;
        }
        @CustomType.Setter
        public Builder interactiveTimeout(GetDatabaseMysqlConfigMysqlInteractiveTimeout interactiveTimeout) {
            if (interactiveTimeout == null) {
              throw new MissingRequiredPropertyException("GetDatabaseMysqlConfigMysql", "interactiveTimeout");
            }
            this.interactiveTimeout = interactiveTimeout;
            return this;
        }
        @CustomType.Setter
        public Builder internalTmpMemStorageEngine(GetDatabaseMysqlConfigMysqlInternalTmpMemStorageEngine internalTmpMemStorageEngine) {
            if (internalTmpMemStorageEngine == null) {
              throw new MissingRequiredPropertyException("GetDatabaseMysqlConfigMysql", "internalTmpMemStorageEngine");
            }
            this.internalTmpMemStorageEngine = internalTmpMemStorageEngine;
            return this;
        }
        @CustomType.Setter
        public Builder maxAllowedPacket(GetDatabaseMysqlConfigMysqlMaxAllowedPacket maxAllowedPacket) {
            if (maxAllowedPacket == null) {
              throw new MissingRequiredPropertyException("GetDatabaseMysqlConfigMysql", "maxAllowedPacket");
            }
            this.maxAllowedPacket = maxAllowedPacket;
            return this;
        }
        @CustomType.Setter
        public Builder maxHeapTableSize(GetDatabaseMysqlConfigMysqlMaxHeapTableSize maxHeapTableSize) {
            if (maxHeapTableSize == null) {
              throw new MissingRequiredPropertyException("GetDatabaseMysqlConfigMysql", "maxHeapTableSize");
            }
            this.maxHeapTableSize = maxHeapTableSize;
            return this;
        }
        @CustomType.Setter
        public Builder netBufferLength(GetDatabaseMysqlConfigMysqlNetBufferLength netBufferLength) {
            if (netBufferLength == null) {
              throw new MissingRequiredPropertyException("GetDatabaseMysqlConfigMysql", "netBufferLength");
            }
            this.netBufferLength = netBufferLength;
            return this;
        }
        @CustomType.Setter
        public Builder netReadTimeout(GetDatabaseMysqlConfigMysqlNetReadTimeout netReadTimeout) {
            if (netReadTimeout == null) {
              throw new MissingRequiredPropertyException("GetDatabaseMysqlConfigMysql", "netReadTimeout");
            }
            this.netReadTimeout = netReadTimeout;
            return this;
        }
        @CustomType.Setter
        public Builder netWriteTimeout(GetDatabaseMysqlConfigMysqlNetWriteTimeout netWriteTimeout) {
            if (netWriteTimeout == null) {
              throw new MissingRequiredPropertyException("GetDatabaseMysqlConfigMysql", "netWriteTimeout");
            }
            this.netWriteTimeout = netWriteTimeout;
            return this;
        }
        @CustomType.Setter
        public Builder sortBufferSize(GetDatabaseMysqlConfigMysqlSortBufferSize sortBufferSize) {
            if (sortBufferSize == null) {
              throw new MissingRequiredPropertyException("GetDatabaseMysqlConfigMysql", "sortBufferSize");
            }
            this.sortBufferSize = sortBufferSize;
            return this;
        }
        @CustomType.Setter
        public Builder sqlMode(GetDatabaseMysqlConfigMysqlSqlMode sqlMode) {
            if (sqlMode == null) {
              throw new MissingRequiredPropertyException("GetDatabaseMysqlConfigMysql", "sqlMode");
            }
            this.sqlMode = sqlMode;
            return this;
        }
        @CustomType.Setter
        public Builder sqlRequirePrimaryKey(GetDatabaseMysqlConfigMysqlSqlRequirePrimaryKey sqlRequirePrimaryKey) {
            if (sqlRequirePrimaryKey == null) {
              throw new MissingRequiredPropertyException("GetDatabaseMysqlConfigMysql", "sqlRequirePrimaryKey");
            }
            this.sqlRequirePrimaryKey = sqlRequirePrimaryKey;
            return this;
        }
        @CustomType.Setter
        public Builder tmpTableSize(GetDatabaseMysqlConfigMysqlTmpTableSize tmpTableSize) {
            if (tmpTableSize == null) {
              throw new MissingRequiredPropertyException("GetDatabaseMysqlConfigMysql", "tmpTableSize");
            }
            this.tmpTableSize = tmpTableSize;
            return this;
        }
        @CustomType.Setter
        public Builder waitTimeout(GetDatabaseMysqlConfigMysqlWaitTimeout waitTimeout) {
            if (waitTimeout == null) {
              throw new MissingRequiredPropertyException("GetDatabaseMysqlConfigMysql", "waitTimeout");
            }
            this.waitTimeout = waitTimeout;
            return this;
        }
        public GetDatabaseMysqlConfigMysql build() {
            final var _resultValue = new GetDatabaseMysqlConfigMysql();
            _resultValue.connectTimeout = connectTimeout;
            _resultValue.defaultTimeZone = defaultTimeZone;
            _resultValue.groupConcatMaxLen = groupConcatMaxLen;
            _resultValue.informationSchemaStatsExpiry = informationSchemaStatsExpiry;
            _resultValue.innodbChangeBufferMaxSize = innodbChangeBufferMaxSize;
            _resultValue.innodbFlushNeighbors = innodbFlushNeighbors;
            _resultValue.innodbFtMinTokenSize = innodbFtMinTokenSize;
            _resultValue.innodbFtServerStopwordTable = innodbFtServerStopwordTable;
            _resultValue.innodbLockWaitTimeout = innodbLockWaitTimeout;
            _resultValue.innodbLogBufferSize = innodbLogBufferSize;
            _resultValue.innodbOnlineAlterLogMaxSize = innodbOnlineAlterLogMaxSize;
            _resultValue.innodbReadIoThreads = innodbReadIoThreads;
            _resultValue.innodbRollbackOnTimeout = innodbRollbackOnTimeout;
            _resultValue.innodbThreadConcurrency = innodbThreadConcurrency;
            _resultValue.innodbWriteIoThreads = innodbWriteIoThreads;
            _resultValue.interactiveTimeout = interactiveTimeout;
            _resultValue.internalTmpMemStorageEngine = internalTmpMemStorageEngine;
            _resultValue.maxAllowedPacket = maxAllowedPacket;
            _resultValue.maxHeapTableSize = maxHeapTableSize;
            _resultValue.netBufferLength = netBufferLength;
            _resultValue.netReadTimeout = netReadTimeout;
            _resultValue.netWriteTimeout = netWriteTimeout;
            _resultValue.sortBufferSize = sortBufferSize;
            _resultValue.sqlMode = sqlMode;
            _resultValue.sqlRequirePrimaryKey = sqlRequirePrimaryKey;
            _resultValue.tmpTableSize = tmpTableSize;
            _resultValue.waitTimeout = waitTimeout;
            return _resultValue;
        }
    }
}
